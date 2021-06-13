package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"linkscape/controllers"
	"linkscape/db_connection"
	"log"
	"net/http"
)

func main(){
		router := mux.NewRouter()
		

	fmt.Println("Starting server on the port 5000")
	fmt.Println(dbconnection.Createconnection())

	log.Fatal(http.ListenAndServe(":5000", router))
}