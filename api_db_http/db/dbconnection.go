package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func DBConnection() {
	DB, errorDb := sql.Open("postgres", "postgres://postgres:psql1330@localhost/labora-proyect-1?sslmode=disable")
	if errorDb != nil {
		log.Fatal(errorDb)
	} else {
		log.Println("Database connection succesful")
	}
	defer DB.Close()

}
