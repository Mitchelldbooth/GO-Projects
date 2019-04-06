package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	db, err := sql.Open("mysql", "root:Mitchell18@tcp(127.0.0.1)/testdb")

	if err != nil {
		panic(err.Error())

	}

	defer db.Close()

	results, err := db.Query("Select personality from users where name = 'Nirosh'")
	if err != nil {
		panic(err.Error())

	}

	for results.Next() {
		var user User

		err = results.Scan(&user.Name)
		if err != nil {
			panic(err.Error())

		}
		fmt.Println(user.Name)
	}

}
