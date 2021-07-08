package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	dbconnection "linkscape/db_connection"
	Models "linkscape/models"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Createresponse struct {
	ID      int64  `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Url     string `json:"url,omitempty"`
	Message string `json:"message,omitempty"`
}

type response struct {
	Status bool `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

func Welcome(w http.ResponseWriter, r *http.Request) {
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
	sqlStatement := `INSERT INTO links (title, link, categoryid, userid) VALUES ($1, $2, $3, $4) RETURNING linkId, link,title`
	// close the db connection
	defer db.Close()
	// the inserted id will store in this id
	var linkId int64
	var linkTitle, linkUrl string

	// execute the sql statement
	// Scan function will save the insert id in the id
	query := db.QueryRow(sqlStatement, link.Title, link.Link, link.CategoryId, link.UserId).Scan(&linkId, &linkUrl, &linkTitle)

	if query != nil {
		log.Fatalf("Unable to execute the query. %v", query)
	}

	fmt.Printf("Inserted a single record %v", linkId)

	res := Createresponse{
		ID:      linkId,
		Title:   linkTitle,
		Url:     linkUrl,
		Message: "Link added successfully",
	}
	// sending back json
	json.NewEncoder(w).Encode(res)
}

//get link by id
func GetLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	// Grabbing the id from the url
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string into int %v", err)
	}
	// create connection
	db := dbconnection.Createconnection()
	// close the db connection
	defer db.Close()
	var link Models.Link
	// sql statement
	sqlStatement := `SELECT * FROM links WHERE linkid = $1`
	// Execute sql statement
	row := db.QueryRow(sqlStatement, id)
	errFromStruct := row.Scan(&link.ID, &link.Title, &link.Link, &link.CategoryId, &link.UserId)
	switch errFromStruct {
	case sql.ErrNoRows:
		w.WriteHeader(http.StatusNotFound)
		res := response{
			Status: false,
			Message: "Sorry, link does not exist",
		}
		json.NewEncoder(w).Encode(res)
	case nil:
		json.NewEncoder(w).Encode(link)
	default:
		log.Fatalf("Unable to scan the row. %v", errFromStruct)
	}

}

func GetAllLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	db := dbconnection.Createconnection()
	var links []Models.Link
	var link Models.Link
	// sql statement
	sqlStatement := `SELECT * FROM links`
	// Execute sql statement
	row, err := db.Query(sqlStatement)
	if err != nil {
		res := response{
			Message: "An error occured",
		}
		json.NewEncoder(w).Encode(res)
	}
	for row.Next() {
		errFromStruct := row.Scan(&link.ID, &link.Title, &link.Link, &link.CategoryId, &link.UserId)
		if errFromStruct != nil {
			res := response{
				Message: "An error occured",
			}
			json.NewEncoder(w).Encode(res)
		}
		links = append(links, link)
	}
	json.NewEncoder(w).Encode(links)
	// close the db connection
	defer db.Close()
}

func UpdateLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("	Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	var link Models.Link

	if err != nil {
		log.Fatalf("Unable to convert string into int %v", err)
	}
	err = json.NewDecoder(r.Body).Decode(&link)

    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }
	db := dbconnection.Createconnection()

	sqlStatement := `UPDATE links SET title=$2, link=$3, categoryid=$4, userId=$5 WHERE linkId=$1`
	// execute the sql statement
	result, err := db.Exec(sqlStatement,id, link.Title, link.Link, link.CategoryId, link.UserId)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}
	if rowsAffected > 0 {
		res := response{
		Status : true,
		Message: "link updated succesfully!",
		}
	json.NewEncoder(w).Encode(res)
	}else{
		w.WriteHeader(http.StatusNotFound)
		res := response{
			Status : false,
			Message: "Could not update, link not found",
	}
	json.NewEncoder(w).Encode(res)

	}	
	
	fmt.Printf("Total rows/record affected %v", rowsAffected)

		// close the db connection
		defer db.Close()
}


func DeleteLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	var link Models.Link

	if err != nil {
		log.Fatalf("Unable to convert string into int %v", err)
	}
	err = json.NewDecoder(r.Body).Decode(&link)

    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }
	db := dbconnection.Createconnection()

	sqlStatement := `DELETE FROM links WHERE linkId=$1`
	// execute the sql statement
	result, err := db.Exec(sqlStatement,id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}
	if rowsAffected > 0 {
		res := response{
		Status : true,
		Message: "link deleted succesfully!",
		}
	json.NewEncoder(w).Encode(res)
	}else{
		w.WriteHeader(http.StatusNotFound)
		res := response{
			Status : false,
			Message: "Could not delete, link does not exist",
	}
	json.NewEncoder(w).Encode(res)

	}	
	
	fmt.Printf("Total rows/record affected %v", rowsAffected)
		// close the db connection
		defer db.Close()

}