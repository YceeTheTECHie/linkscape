package dbconnection

import (
	"database/sql"
	"log"
	"os"
	"fmt"
	_"github.com/lib/pq"
	"github.com/joho/godotenv"
)

func Createconnection() *sql.DB {
    // load .env file
	err := godotenv.Load(".env")
	

    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    // Open the connection
    db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

    if err != nil {
        panic(err)
    }

    // check the connection
    err = db.Ping()

    if err != nil {
        panic(err)
    }

    fmt.Println("Database successfully connected!")
    // return the connection
    return db
}

