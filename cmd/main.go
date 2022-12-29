package main

import (
	"fmt"
	"log"
	"net/http"
	"test-api/pkg/controllers"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/Get", controllers.GetAll).Methods("GET")
	router.HandleFunc("/Get/{uuid}", controllers.Get).Methods("GET")
	router.HandleFunc("/Add", controllers.AddBook).Methods("POST")
	router.HandleFunc("/Update", controllers.Update).Methods("PATCH")
	router.HandleFunc("/homelink", homeLink).Methods("GET")
	router.HandleFunc("/create-table", controllers.CreateDynamoDbTable).Methods("POST")
	router.HandleFunc("/get-tables", controllers.ListDynamoDbTables).Methods("GET")

	log.Fatal(http.ListenAndServe(":8082", router))
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Welcome home!")
}

// func listDynamoDbTables(sess *session.Session) {
// 	svc := dynamodb.New(sess)

// 	// create the input configuration instance
// 	input := &dynamodb.ListTablesInput{}
// 	for {
// 		// Get the list of tables
// 		result, err := svc.ListTables(input)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			if aerr, ok := err.(awserr.Error); ok {
// 				switch aerr.Code() {
// 				case dynamodb.ErrCodeInternalServerError:
// 					fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
// 				default:
// 					fmt.Println(aerr.Error())
// 				}
// 			} else {
// 				// Print the error, cast err to awserr.Error to get the Code and
// 				// Message from an error.
// 				fmt.Println(err.Error())
// 			}
// 			return
// 		}

// 		for _, n := range result.TableNames {
// 			fmt.Println(*n)
// 		}

// 		// assign the last read tablename as the start for our next call to the ListTables function
// 		// the maximum number of table names returned in a call is 100 (default), which requires us to make
// 		// multiple calls to the ListTables function to retrieve all table names
// 		input.ExclusiveStartTableName = result.LastEvaluatedTableName

// 		if result.LastEvaluatedTableName == nil {
// 			break
// 		}
// 	}
// }

// func createDynamoDbTable(sess *session.Session) {
// 	svc := dynamodb.New(sess)

// 	// Create table Movies
// 	tableName := "Books"

// 	/*
// 			type Book struct {
// 			ID          int    `json:"id"`
// 			uuid		string	`json:"uuid"`
// 			Name        string `json:"name"`
// 			Author      string `json:"author"`
// 			Publication string `json:"publication"`
// 		}
// 	*/

// 	input := &dynamodb.CreateTableInput{
// 		AttributeDefinitions: []*dynamodb.AttributeDefinition{
// 			{
// 				AttributeName: aws.String("uuid"),
// 				AttributeType: aws.String("S"),
// 			},
// 			{
// 				AttributeName: aws.String("name"),
// 				AttributeType: aws.String("S"),
// 			},
// 		},
// 		KeySchema: []*dynamodb.KeySchemaElement{
// 			{
// 				AttributeName: aws.String("uuid"),
// 				KeyType:       aws.String("HASH"),
// 			},
// 			{
// 				AttributeName: aws.String("name"),
// 				KeyType:       aws.String("RANGE"),
// 			},
// 		},
// 		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
// 			ReadCapacityUnits:  aws.Int64(10),
// 			WriteCapacityUnits: aws.Int64(10),
// 		},
// 		TableName: aws.String(tableName),
// 	}

// 	_, err := svc.CreateTable(input)
// 	if err != nil {
// 		log.Fatalf("Got error calling CreateTable: %s", err)
// 	}

// 	fmt.Println("Created the table", tableName)
// }
