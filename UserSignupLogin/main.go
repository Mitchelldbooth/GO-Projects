package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {
	http.HandleFunc("/login", Login)
	http.HandleFunc("/signup", Signup)
	// initialize our database connection
	initDB()
	// start the server on port 8000 local 
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func initDB() {
	var err error

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	password := os.Getenv("db_pass")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_IP")

	// Connect to db
	_, err = sql.Open("mysql", fmt.Sprintf("root:%s@tcp(%s)/%s", password, dbHost, dbName))
	if err != nil {
		panic(err)
	}
}
