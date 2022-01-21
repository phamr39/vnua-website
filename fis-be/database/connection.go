package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	user     = "root"
	password = "12345678"
	host     = "103.170.122.117:3305"
	database = "admin_vnuapharmacy"
)

func Connect() (db *sql.DB) {
	dsn := user + ":" + password + "@tcp(" + host + ")/" + database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Could not connect to database")
	}
	return db
}
