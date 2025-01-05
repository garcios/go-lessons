package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "events.db")
	if err != nil {
		panic(fmt.Sprintf("Could not connect to database: %v", err))
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable()
	createEventsTable()
	createRegistrationsTable()
}

func createUsersTable() {
	_, err := DB.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
`)
	if err != nil {
		panic(fmt.Sprintf("Could not create users table: %v", err))
	}
}

func createEventsTable() {
	_, err := DB.Exec(`
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER,
	    FOREIGN KEY(id) REFERENCES users(id)                              
	)
`)
	if err != nil {
		panic(fmt.Sprintf("Could not create events table: %v", err))
	}
}

func createRegistrationsTable() {
	_, err := DB.Exec(`
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
`)
	if err != nil {
		panic(fmt.Sprintf("Could not create registrations table: %v", err))
	}
}
