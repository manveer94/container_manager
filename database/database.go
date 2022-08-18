package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// Initializes the database
func Init() {
	if db == nil {
		var err error
		db, err = sql.Open("sqlite3", "aqua.db")

		if err != nil {
			log.Fatalf("error while connecting to database : %s", err.Error())
		}
		db.SetMaxOpenConns(5)
	}
}

// Returns the instance of the database
func Get() *sql.DB {
	return db
}
