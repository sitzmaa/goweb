package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// InitializeDatabase sets up the database connection and creates tables if they don't exist
func InitializeDatabase() {
	var err error
	db, err = sql.Open("sqlite3", "./goweb-server.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create a table for storing API logs, for example
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS logs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		action TEXT,
		status TEXT,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Fatal(err)
	}
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	return db
}
