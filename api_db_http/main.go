package main

import (
	"net/http"

	"github.com/ThaliaAC/labora-api/api_db_http/db"
	"github.com/ThaliaAC/labora-api/api_db_http/handler"
	"github.com/gorilla/mux"
)

func main() {
	db.ConnectToDb(db.Db)
	router := mux.NewRouter()

	router.HandleFunc("/", handler.IndexRoute).Methods("GET")
	router.HandleFunc("/items", handler.GetUsersHandler).Methods("GET")
	router.HandleFunc("/item/{id}", handler.GetUserByIdHandler).Methods("GET")
	router.HandleFunc("/items/{name}", handler.GetUserByNameHandler).Methods("GET")
	//router.HandleFunc("/items", handler.CreateItemHandler).Methods("POST")
	router.HandleFunc("/items/{id}", handler.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/items/{id}", handler.DeleteUserHandler).Methods("DELETE")
	router.HandleFunc("/itemdetails", handler.GetUserDetailsHandler).Methods("GET")
	http.ListenAndServe(":9000", router)
}
