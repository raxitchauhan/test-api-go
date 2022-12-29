package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"test-api/pkg/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
)

func WriteRespose(w http.ResponseWriter, v any, statusCode int) {
	res, err := json.Marshal(v)

	if err != nil {
		ThrowError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(res)
}

func ThrowError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func SetAwsSession() (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String("us-west-2"),
		Credentials:      credentials.NewStaticCredentials("test", "test", ""),
		S3ForcePathStyle: aws.Bool(true),
		Endpoint:         aws.String("http://localhost:4566"),
	})

	return sess, err
}

// AddTableItem adds an item to an Amazon DynamoDB table
// Inputs:
//
//	sess is the current session, which provides configuration for the SDK's service clients
//	year is the year when the movie was released
//	table is the name of the table
//	title is the movie title
//	plot is a summary of the plot of the movie
//	rating is the movie rating, from 0.0 to 10.0
//
// Output:
//
//	If success, nil
//	Otherwise, an error from the call to PutItem
func AddTableItem(svc dynamodbiface.DynamoDBAPI, book *models.Book, table *string) error {

	av, err := dynamodbattribute.MarshalMap(book)
	// snippet-end:[dynamodb.go.create_new_item.assign_struct]
	if err != nil {
		return err
	}

	// snippet-start:[dynamodb.go.create_new_item.call]
	_, err = svc.PutItem(&dynamodb.PutItemInput{
		Item:      av,
		TableName: table,
	})
	// snippet-end:[dynamodb.go.create_new_item.call]
	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}

func PublishToSns(svc snsiface.SNSAPI, msg, topicARN *string) {

	result, err := svc.Publish(&sns.PublishInput{
		Message:  msg,
		TopicArn: topicARN,
	})

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(*result.MessageId)
}
