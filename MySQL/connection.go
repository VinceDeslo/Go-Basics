package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Entering MySQL connection program.")

	// Open the database connection
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/db")

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Opened db successfully.")
	defer db.Close()

	// Create a table for the entries
	_, err = db.Exec("CREATE Table members(id int NOT NULL AUTO_INCREMENT, first_name varchar(50), last_name varchar(50), PRIMARY KEY(id));")

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Created members table successfully.")

	// Insert a new entry into the mysql database
	insert, err := db.Query("INSERT INTO members (first_name, last_name) VALUES ('Vincent','Desloover')")

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Inserted a member entry successfully.")

	defer insert.Close()
}
