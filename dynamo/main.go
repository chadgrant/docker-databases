package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func main() {
	log.Println("starting populator ...")

	ep := os.Getenv("DYNAMO_ENDPOINT")
	if len(ep) == 0 {
		ep = "http://localhost:8000"
	}

	dyn := dynamodb.New(session.Must(session.NewSession()), &aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("key", "secret", ""),
		Endpoint:    aws.String(ep),
	})

	var dir *os.File
	var err error

	_, err = os.Stat("/data")
	if !os.IsNotExist(err) {
		dir, err = os.Open("/data")
		if err != nil {
			panic(fmt.Errorf("could not open /data %v", err))
		}
	} else {
		_, err = os.Stat("sample/data")
		if !os.IsNotExist(err) {
			dir, err = os.Open("sample/data")
			if err != nil {
				panic(fmt.Errorf("could not open sample/data %v", err))
			}
		}
	}

	dirs, err := dir.Readdirnames(0)
	if err != nil {
		panic(fmt.Errorf("could not read dirs in %s %v", dir.Name(), err))
	}

	for _, d := range dirs {
		log.Printf("checking directory %s", d)
		file := path.Join(dir.Name(), d, "schema.json")
		_, err := os.Stat(file)
		if !os.IsNotExist(err) {
			log.Printf("creating schema %s\n", d)
			if err := createTable(dyn, true, file); err != nil {
				log.Printf("couldn't create table %s %v\n", d, err)
			}
		}

		file = path.Join(dir.Name(), d, "data.json")
		_, err = os.Stat(file)
		if !os.IsNotExist(err) {
			log.Printf("populating table %s\n", d)
			if err = loadDataFromFile(dyn, d, file); err != nil {
				panic(err)
			}
		}
	}
}

func loadDataFromFile(dynamo *dynamodb.DynamoDB, table, file string) error {

	if !isTableActive(dynamo, table, 3*time.Second) {
		return fmt.Errorf("timeout checking table state: %s", table)
	} else {
		log.Printf("importing data %s from %s\n", table, file)
		bs, err := ioutil.ReadFile(file)
		if err != nil {
			return fmt.Errorf("reading data from disk %s %v", file, err)
		}
		var rows []json.RawMessage
		if err := json.Unmarshal(bs, &rows); err != nil {
			return fmt.Errorf("could not unmarshal data %s %v", table, err)
		}

		for _, r := range rows {
			var row map[string]interface{}
			if err := json.Unmarshal(r, &row); err != nil {
				return fmt.Errorf("could not unmarshal row %s %v", table, err)
			}
			av, err := dynamodbattribute.MarshalMap(row)
			if err != nil {
				return fmt.Errorf("error marshalling %v", err)
			}
			_, err = dynamo.PutItem(&dynamodb.PutItemInput{
				Item:      av,
				TableName: aws.String(table),
			})
			if err != nil {
				return fmt.Errorf("error putitem %s %v", table, err)
			}
		}
		log.Printf("added %d rows to %s\n", len(rows), table)
	}

	return nil
}

func createTable(dynamo *dynamodb.DynamoDB, deleteTable bool, file string) error {
	t, err := loadTableSchema(file)
	if err != nil {
		return err
	}

	_, err = dynamo.CreateTable(t)
	return err
}

func isTableActive(dynamo *dynamodb.DynamoDB, table string, timeout time.Duration) bool {
	tick := time.NewTicker(500 * time.Millisecond)
	timeoutC := time.After(timeout)
	defer tick.Stop()

	for {
		select {
		case <-timeoutC:
			return false

		case <-tick.C:
			if resp, err := dynamo.DescribeTable(&dynamodb.DescribeTableInput{TableName: aws.String(table)}); err == nil {
				if *resp.Table.TableStatus == "ACTIVE" {
					return true
				}
			}
		}
	}
}

func loadTableSchema(file string) (*dynamodb.CreateTableInput, error) {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("reading schema file %s %v", file, err)
	}

	t := &dynamodb.CreateTableInput{}
	if err = json.Unmarshal(bs, t); err != nil {
		return nil, fmt.Errorf("unmarshaling schema %s %v", file, err)
	}

	return t, nil
}
