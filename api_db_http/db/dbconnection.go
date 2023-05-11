package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	dbName   = "labora-proyect-1"
	user     = "postgres"
	password = "psql1330"
)

var Db *sql.DB

func ConnectToDb(database *sql.DB) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	dbConn, dbErr := sql.Open("postgres", psqlInfo)
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	defer dbConn.Close()
	return dbConn, dbErr
}
