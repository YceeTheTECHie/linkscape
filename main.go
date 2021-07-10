package main

import (
	"fmt"
	"linkscape/controllers"
	"linkscape/db_connection"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

	
func main(){
	router := mux.NewRouter()
	router.HandleFunc("/api/newlink", controllers.CreateLink).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/links", controllers.GetAllLink).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/link/{id}", controllers.GetLink).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/link/{id}", controllers.UpdateLink).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/link/{id}", controllers.DeleteLink).Methods("DELETE", "OPTIONS")


	router.HandleFunc("/", controllers.Welcome).Methods("POST", "OPTIONS")
	port := os.Getenv("PORT")
	fmt.Println("Starting server on the port 5000")
	fmt.Println(dbconnection.Createconnection())
	log.Fatal(http.ListenAndServe(port, router))
}