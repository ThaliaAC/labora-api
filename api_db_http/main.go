package main

import (
	"log"
	"net/http"

	"github.com/ThaliaAC/labora-api/api_db_http/db"
	"github.com/ThaliaAC/labora-api/api_db_http/handler"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	db.ConnectToDb()
	router := mux.NewRouter()

	router.HandleFunc("/", handler.IndexRoute).Methods("GET")
	router.HandleFunc("/items", handler.GetItemsHandler).Methods("GET")
	router.HandleFunc("/item/{id}", handler.GetItemsByIdHandler).Methods("GET")
	router.HandleFunc("/items/{customerName}", handler.GetItemsByNameHandler).Methods("GET")
	router.HandleFunc("/items", handler.CreateItemHandler).Methods("POST")
	router.HandleFunc("/item/{id}", handler.UpdateItemHandler).Methods("PUT")
	router.HandleFunc("/item/{id}", handler.DeleteItemHandler).Methods("DELETE")

	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:9000/"},
		AllowedMethods: []string{"PUT"},
	})
	handler := corsOptions.Handler(router)

	log.Fatal(http.ListenAndServe(":9000", handler))

}
