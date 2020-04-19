package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func main() {
	log.Println("starting populator ...")
	defer log.Println("populator complete ...")

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
			} else {
				loadAllFiles(dyn, d, path.Join(dir.Name(), d))
			}
		}
	}
}

func loadAllFiles(dyn *dynamodb.DynamoDB, table, dir string) {
	fi, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, f := range fi {
		if f.IsDir() || strings.EqualFold(f.Name(), "schema.json") {
			continue
		}
		if err = loadDataFromFile(dyn, table, path.Join(dir, f.Name())); err != nil {
			panic(err)
		}
	}
}

func loadDataFromFile(dynamo *dynamodb.DynamoDB, table, file string) error {
	if !isTableActive(dynamo, table, 3*time.Second) {
		return fmt.Errorf("timeout checking table state: %s", table)
	}

	readByLines, rerr := canReadByLine(file)
	if rerr != nil {
		return fmt.Errorf("Reading file %s %v", file, rerr)
	}

	var count int64
	var err error
	log.Printf("importing data %s from %s\n", table, file)
	if readByLines {
		count, err = insertFileByLine(dynamo, table, file)
		if err != nil {
			return fmt.Errorf("inserting file line by line %v", err)
		}
	} else {
		count, err = insertEntireFile(dynamo, table, file)
		if err != nil {
			return fmt.Errorf("inserting entire file %v", err)
		}
	}
	log.Printf("added %d rows to %s\n", count, table)

	return nil
}

func insertFileByLine(dynamo *dynamodb.DynamoDB, table, file string) (int64, error) {
	f, err := os.OpenFile(file, os.O_RDWR, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	var counter int64
	r := bufio.NewScanner(f)
	for r.Scan() {
		t := strings.Trim(r.Text(), " ")
		t = strings.TrimRight(t, ",")
		t = strings.TrimRight(t, "]")
		t = strings.TrimLeft(t, "[")
		t = strings.Trim(t, " ")

		if len(t) == 0 {
			continue
		}

		if strings.HasPrefix(t, "{") && strings.HasSuffix(t, "}") {
			counter++

			var row json.RawMessage
			if err := json.Unmarshal([]byte(t), &row); err != nil {
				return 0, fmt.Errorf("could not unmarshal data %s %v", table, err)
			}

			if err := writeRow(dynamo, table, row); err != nil {
				return counter, err
			}
		}
	}
	return counter, nil
}

func insertEntireFile(dynamo *dynamodb.DynamoDB, table, file string) (int64, error) {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		return 0, fmt.Errorf("reading data from disk %s %v", file, err)
	}
	var rows []json.RawMessage
	if err := json.Unmarshal(bs, &rows); err != nil {
		return 0, fmt.Errorf("could not unmarshal data %s %v", table, err)
	}
	for _, row := range rows {
		if err := writeRow(dynamo, table, row); err != nil {
			return 0, err
		}
	}

	return int64(len(rows)), nil
}

func writeRow(dynamo *dynamodb.DynamoDB, table string, r []byte) error {
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

func canReadByLine(filename string) (bool, error) {
	f, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return false, err
	}
	defer f.Close()

	r := bufio.NewScanner(f)
	for r.Scan() {
		t := strings.Trim(r.Text(), " ")
		t = strings.TrimRight(t, ",")
		t = strings.TrimRight(t, "]")
		t = strings.TrimLeft(t, "[")
		t = strings.Trim(t, " ")

		if len(t) == 0 {
			continue
		}

		if strings.HasPrefix(t, "{") && strings.HasSuffix(t, "}") {
			continue
		}

		return false, nil
	}

	return true, nil
}
