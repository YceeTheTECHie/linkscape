package controllers

import (
	"encoding/json"
	"fmt"
	dbconnection "linkscape/db_connection"
	Models "linkscape/models"
	"log"
	"net/http"
	"strconv"
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Createresponse struct {
	ID      int64  `json:"id,omitempty"`
	Title      string  `json:"title,omitempty"`
	Url 	string `json:"url,omitempty"`
    Message string `json:"message,omitempty"`
}

type response struct {
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
	w.Header().Set("Content-Type", "application/json")
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
	sqlStatement := `INSERT INTO links (title, link, categoryid, userid) VALUES ($1, $2, $3, $4) RETURNING linkId, link,title`

    // the inserted id will store in this id
	var linkId int64
	var linkTitle,  linkUrl string

    // execute the sql statement
    // Scan function will save the insert id in the id
    query:= db.QueryRow(sqlStatement, link.Title, link.Link, link.CategoryId, link.UserId).Scan(&linkId, &linkUrl,&linkTitle)

    if query != nil {
        log.Fatalf("Unable to execute the query. %v", query)
    }

	fmt.Printf("Inserted a single record %v", linkId)
	
	res := Createresponse{
		ID:      linkId,
		Title : linkTitle,
		Url: linkUrl,
        Message: "Link added successfully",
    }
// sending back json
	json.NewEncoder(w).Encode(res)
}

//get link by id
func GetLink(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")	
	// Grabbing the id from the url
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string into int %v",err)
	}
	// create connection
	db := dbconnection.Createconnection()
	// close the db connection
	defer db.Close()	
	var link Models.Link
	// sql statement
	sqlStatement := `SELECT * FROM links WHERE linkid = $1`
	// Execute sql statement
	row := db.QueryRow(sqlStatement,id)
	errFromStruct := row.Scan(&link.ID,&link.Title,&link.Link,&link.CategoryId,&link.UserId)
	switch errFromStruct {
	case sql.ErrNoRows:
		res := response{
			Message: "User created successfully",
		   }
	   
		json.NewEncoder(w).Encode(res)
    case nil:
           json.NewEncoder(w).Encode(link)
    default:
        log.Fatalf("Unable to scan the row. %v", errFromStruct)
    }

	







}
