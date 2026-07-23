package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

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
	// you can use _ = os.Remove(dbName) to clear db on every run for practice...but never in production !

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		fmt.Println("closing database connection")
		if err := db.Close(); err != nil {
			log.Printf("error closing database connection: %v", err)
		}
	}()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("database conection established")

	_, err = db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("table was created")
}
