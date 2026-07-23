package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// database schema used to initialize the users table
var schema = `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		hashed_password BLOB NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	`

func main() {
	dbName := "data.db"
	_ = os.Remove(dbName) //clears db on every run....for practice purposes here...never use in production !

	db, err := sql.Open("sqlite3", dbName) //open a database handle using the SQLite driver
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		fmt.Println("closing database connection")
		if err := db.Close(); err != nil {
			log.Printf("error closing database connection: %v", err)
		}
	}()

	err = db.Ping() //verify that the database connection is ready
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("database conection established")

	_, err = db.Exec(schema) //execute the schema to create the users table
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("table was created")
}
