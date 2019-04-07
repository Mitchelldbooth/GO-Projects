package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

type User struct {
	uuid       string
	customerID string // stripe
	email      string `json:"email", db:"email"`
	fullname   string `json:"username", db:"username"`
	address    string `json:"address", db:"address"`
	phone      string `json:"phone", db:"phone"`
	lastLogin  int64
	services   []int32
}

// Signup the
func Signup(w http.ResponseWriter, r *http.Request) {
	// Parse and decode the request body into a new `Credentials` instance

	fmt.Println("sucesssssssasaas")
	credentials := &User{}
	err := json.NewDecoder(r.Body).Decode(credentials)
	if err != nil {
		// If there is something wrong badrequest error 400
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// insert the user into the database
	if _, err = db.Query("insert into users values ($1, $2)",
		credentials.email, credentials.fullname); err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// We reach this point if the email was correctly stored in the database, and the default status of 200 is sent back
}

// Login the user
func Login(w http.ResponseWriter, r *http.Request) {

	// Parse and decode the request body into a new `User` instance
	credentials := &User{}
	err := json.NewDecoder(r.Body).Decode(credentials)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Get the existing entry present in the database for the given username
	result := db.QueryRow("select email from users where email=$1", credentials.email)
	if err != nil {
		// If there is an issue with the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// We create another instance of `Credentials` to store the credentials we get from the database
	storedCreds := &User{}

	err = result.Scan(&storedCreds.email)
	if err != nil {
		// If email not there "Unauthorized"(401)
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// If the error is of any other type, send a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// If we reach this point, send the email
	// The default 200 status is sent
}
