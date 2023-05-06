package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Item struct {
	ID           string    `json:"id"`
	CustomerName string    `json:"customerName"`
	OrderDate    time.Time `json:"orderDate"`
	Product      string    `json:"product"`
	Quantity     int       `json:"quantity"`
	Price        float32   `json:"price"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/items", getItems).Methods("GET")
	http.ListenAndServe(":8000", router)

}

// Función para obtener todos los elementos
func getItemsDB() []Item {
	db, errDb := sql.Open("postgres", "postgres://postgres:psql1330@localhost/labora-proyect-1?sslmode=disable")
	if errDb != nil {
		log.Fatal(errDb)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM items")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var items []Item = make([]Item, 0)

	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return items
}

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items := getItemsDB()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}
