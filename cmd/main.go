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
