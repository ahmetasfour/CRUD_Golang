// database.go
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDatabase() {
	var err error
	db, err = sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	fmt.Println("Database connected successfully")
	createTable()
}

func createTable() {
	createUserTableSQL := `CREATE TABLE IF NOT EXISTS users (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "name" TEXT,
        "email" TEXT
    );`

	statement, err := db.Prepare(createUserTableSQL)
	if err != nil {
		log.Fatal("Error preparing table creation: ", err)
	}
	defer statement.Close()

	_, err = statement.Exec()
	if err != nil {
		log.Fatal("Error executing table creation: ", err)
	}

	fmt.Println("Users table ensured")
}
