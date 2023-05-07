package main

import (
	"net/http"

	"github.com/ThaliaAC/labora-api/api_db_http/db"
	"github.com/ThaliaAC/labora-api/api_db_http/handler"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()
	router := mux.NewRouter()

	router.HandleFunc("/items", handler.GetUser).Methods("GET")
	router.HandleFunc("/item/{id}", handler.GetUserID).Methods("GET")

	http.ListenAndServe(":9000", router)
}
