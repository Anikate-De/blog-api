package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Connect to the SQLite DB
func Connect() {
	DB, err := sql.Open("sqlite3", "./blog.db")

	if err != nil {
		panic(err)
	}

	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(10)

	log.Printf("Connected to the Database")
}
