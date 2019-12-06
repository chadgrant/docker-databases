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

	enumFiles("/schema", func(dir, file string) error {
		log.Printf("creating schema %s\n", file)
		if err := createTable(dyn, true, path.Join(dir, file)); err != nil {
			log.Printf("couldn't create table %s %v\n", path.Join(dir, file), err)
		}
		return nil
	})

	enumFiles("/data", func(dir, file string) error {
		t, err := loadTableSchema(path.Join("/schema", file))
		if err != nil {
			log.Println(err)
		}
		if !isTableActive(dyn, *t.TableName, 3*time.Second) {
			return fmt.Errorf("timeout checking table state: %s", file)
		} else {
			log.Printf("importing data %s\n", path.Join(dir, file))
			bs, err := ioutil.ReadFile(path.Join(dir, file))
			if err != nil {
				return fmt.Errorf("reading data from disk %s %v", path.Join(dir, file), err)
			}
			var rows []json.RawMessage
			if err := json.Unmarshal(bs, &rows); err != nil {
				return fmt.Errorf("could not unmarshal data %s %v", *t.TableName, err)
			}

			for _, r := range rows {
				var row map[string]interface{}
				if err := json.Unmarshal(r, &row); err != nil {
					return fmt.Errorf("could not unmarshal row %s %v", *t.TableName, err)
				}
				av, err := dynamodbattribute.MarshalMap(row)
				if err != nil {
					return fmt.Errorf("error marshalling %v", err)
				}
				_, err = dyn.PutItem(&dynamodb.PutItemInput{
					Item:      av,
					TableName: aws.String(*t.TableName),
				})
				if err != nil {
					return fmt.Errorf("error putitem %s %v", *t.TableName, err)
				}
			}
			log.Printf("added %d rows to %s\n", len(rows), *t.TableName)
		}

		return nil
	})
}

func enumFiles(dir string, fun func(string, string) error) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("reading data files %v\n", err)
	}

	for _, f := range files {
		if err := fun(dir, f.Name()); err != nil {
			return err
		}
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
