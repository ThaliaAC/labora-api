package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Item struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
}

var items []Item

func main() {
	for i := 1; i <= 10; i++ {
		items = append(items, Item{ID: fmt.Sprintf("%d", i), Name: fmt.Sprintf("Item %d", i)})
	}

	router := mux.NewRouter()
	router.HandleFunc("/", indexRoute).Methods("GET")
	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items/{id}", getItemID).Methods("GET")
	router.HandleFunc("/items/{name}", getItemName).Methods("GET")
	router.HandleFunc("/items", createItem).Methods("POST")
	router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

	var portNumber int = 3000
	fmt.Println("Listen in port ", portNumber)
	err := http.ListenAndServe(":"+strconv.Itoa(portNumber), router)
	if err != nil {
		fmt.Println(err)
	}

}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my first API")
}

// Función para obtener todos los elementos
func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// Función para obtener un elemento específico
func getItemID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(r)

	for _, item := range items {
		if item.ID == parameters["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode("Item not found")
}

// Función para buscar un elemento por nombre
func getItemName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(r)

	for _, item := range items {
		if strings.EqualFold(item.Name, parameters["name"]) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode("Item not found")
}

// Función para crear un nuevo elemento
func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Recibo la información
	var newItem Item
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Item")
		return
	}
	//Asigno esa información a la variable newItem
	json.Unmarshal(reqBody, &newItem)

	//Automatizo incremento de ID
	id := len(items) + 1
	newItem.ID = "item" + strconv.Itoa(id)
	//Guardo newItem en la lista de items
	items = append(items, newItem)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
	fmt.Println("Item successfully created")
}

// Función para actualizar un elemento existente
func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updatedItem Item
	parameters := mux.Vars(r)
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter Valid Data")
		return
	}

	json.Unmarshal(reqBody, &updatedItem)

	for index, item := range items {
		if item.ID == parameters["id"] {
			items = append(items[:index], items[index+1:]...)
			updatedItem.ID = parameters["id"]
			items = append(items, updatedItem)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedItem)
	fmt.Println("The item has been updated successfully")
}

// Función para eliminar un elemento
func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(r)
	for index, item := range items {
		if item.ID == parameters["id"] {
			items = append(items[:index], items[index+1:]...)
			fmt.Println(w, "The item has been remove succesfully")
			return
		}
	}
	_, err := strconv.Atoi(parameters["id"])
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}
}
