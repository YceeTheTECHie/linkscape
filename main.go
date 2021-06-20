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
	router.HandleFunc("/api/newlink", controllers.CreateLink).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/links", controllers.GetAllLink).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/link/{id}", controllers.GetLink).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/link/{id}", controllers.UpdateLink).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/link/{id}", controllers.DeleteLink).Methods("DELETE", "OPTIONS")


	router.HandleFunc("/", controllers.Welcome).Methods("POST", "OPTIONS")
	fmt.Println("Starting server on the port 5000")
	fmt.Println(dbconnection.Createconnection())
	log.Fatal(http.ListenAndServe(":5000", router))
}