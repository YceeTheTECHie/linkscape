package controllers

import (
	"encoding/json"
	dbconnection "linkscape/db_connection"
	Models "linkscape/models"
	"log"
	"net/http"
	"fmt"
)

type response struct {
    ID      int64  `json:"id,omitempty"`
    Message string `json:"message,omitempty"`
}


func Welcome(w http.ResponseWriter, r *http.Request){
	res := response{
     Message: "User created successfully",
    }

    // send the response
    json.NewEncoder(w).Encode(res)
}
func CreateLink(w http.ResponseWriter, r *http.Request) {
	//setting up headers
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	//creating an empty link of type Link
	var link Models.Link

	err := json.NewDecoder(r.Body).Decode(&link)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}
	
	// create the postgres db connection
	  db := dbconnection.Createconnection()

	  // close the db connection
	  defer db.Close()	

	  sqlStatement := `INSERT INTO links (title, link, categoryid, userid) VALUES ($1, $2, $3, $4) RETURNING linkid`

    // the inserted id will store in this id
    var linkId int64

    // execute the sql statement
    // Scan function will save the insert id in the id
    query:= db.QueryRow(sqlStatement, link.Title, link.Link, link.CategoryId, link.UserId).Scan(&linkId)

    if query != nil {
        log.Fatalf("Unable to execute the query. %v", query)
    }

	fmt.Printf("Inserted a single record %v", linkId)
	
	res := response{
        ID:      linkId,
        Message: "User created successfully",
    }

	json.NewEncoder(w).Encode(res)

    // return the inserted id


}
