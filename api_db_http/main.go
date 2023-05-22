package main

import (
	"log"
	"net/http"

	"github.com/ThaliaAC/labora-api/api_db_http/db"
	"github.com/ThaliaAC/labora-api/api_db_http/handler"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	db.ConnectToDb()
	router := mux.NewRouter()

	router.HandleFunc("/", handler.IndexRoute).Methods("GET")
	router.HandleFunc("/items", handler.GetItemsHandler).Methods("GET")
	router.HandleFunc("/item/{id}", handler.GetItemsByIdHandler).Methods("GET")
	router.HandleFunc("/items/{customerName}", handler.GetItemsByNameHandler).Methods("GET")
	router.HandleFunc("/items", handler.CreateItemHandler).Methods("POST")
	router.HandleFunc("/items/{id}", handler.UpdateItemHandler).Methods("PUT")
	router.HandleFunc("/items/{id}", handler.DeleteItemHandler).Methods("DELETE")

	corsOption1 := handlers.AllowedOrigins([]string{"*"})
	corsOption2 := handlers.AllowedMethods([]string{"POST"})

	handler := handlers.CORS(corsOption1, corsOption2)(router)

	log.Fatal(http.ListenAndServe(":9000", handler))
}
