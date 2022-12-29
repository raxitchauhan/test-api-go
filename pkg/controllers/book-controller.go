package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"test-api/pkg/models"
	"test-api/pkg/utils"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	books, _ := models.GetAll()

	utils.WriteRespose(w, books, http.StatusOK)
}

func Get(w http.ResponseWriter, r *http.Request) {

	uuid := mux.Vars(r)["uuid"]

	book, err := models.Get(uuid)
	if err != nil {
		log.Error(err.Error())
		utils.WriteRespose(w, nil, http.StatusNotFound)
		return
	}

	utils.WriteRespose(w, book, http.StatusOK)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var book *models.Book
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ThrowError(w, err)
	}

	json.Unmarshal(reqBody, &book)

	models.Create(book)

	go func() {
		sess, _ := utils.SetAwsSession()
		// dynamoClient := dynamodb.New(sess)
		snsClient := sns.New(sess)
		// models.AddTableItem(dynamoClient, book, aws.String("Books"))
		// models.PublishToSns(snsClient, aws.String(string(reqBody[:])), aws.String("arn:aws:sns:us-west-2:000000000000:first-proj-sns"))
		b, _ := json.Marshal(book)
		utils.PublishToSns(snsClient, aws.String(string(b[:])), aws.String("arn:aws:sns:us-west-2:000000000000:first-proj-sns"))
	}()

	utils.WriteRespose(w, book, http.StatusCreated)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ThrowError(w, err)
	}

	json.Unmarshal(reqBody, &book)

	book, err = models.Update(book)

	if err != nil {
		log.Error(err.Error())
		utils.WriteRespose(w, nil, http.StatusNotFound)
		return
	}

	utils.WriteRespose(w, book, http.StatusOK)
}

func CreateDynamoDbTable(w http.ResponseWriter, r *http.Request) {
	sess, _ := utils.SetAwsSession()
	svc := dynamodb.New(sess)

	var tableN *models.Table
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ThrowError(w, err)
	}

	json.Unmarshal(reqBody, &tableN)
	tableName := tableN.Name

	/*
			type Book struct {
			ID          int    `json:"id"`
			uuid		string	`json:"uuid"`
			Name        string `json:"name"`
			Author      string `json:"author"`
			Publication string `json:"publication"`
		}
	*/

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("uuid"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("name"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("uuid"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("name"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}

	_, err = svc.CreateTable(input)
	if err != nil {
		log.Errorf("Got error calling CreateTable: %s", err)
		utils.ThrowError(w, err)
		return
	}

	fmt.Println("Created the table", tableName)
	utils.WriteRespose(w, tableName, http.StatusOK)
}

func ListDynamoDbTables(w http.ResponseWriter, r *http.Request) {
	sess, _ := utils.SetAwsSession()
	svc := dynamodb.New(sess)
	var tables []string
	// create the input configuration instance
	input := &dynamodb.ListTablesInput{}
	for {
		// Get the list of tables
		result, err := svc.ListTables(input)
		if err != nil {
			fmt.Println(err.Error())
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case dynamodb.ErrCodeInternalServerError:
					fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return
		}

		for _, n := range result.TableNames {
			// fmt.Println(*n)
			tables = append(tables, *n)
		}

		// assign the last read tablename as the start for our next call to the ListTables function
		// the maximum number of table names returned in a call is 100 (default), which requires us to make
		// multiple calls to the ListTables function to retrieve all table names
		input.ExclusiveStartTableName = result.LastEvaluatedTableName

		if result.LastEvaluatedTableName == nil {
			break
		}
	}

	utils.WriteRespose(w, tables, http.StatusOK)
}
