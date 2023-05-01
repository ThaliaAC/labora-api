package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func raiz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Estás en la ruta raiz")
}
func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pageUser := r.URL.Query().Get("page")
	itemsUser := r.URL.Query().Get("itemsPerPage")
	page, err := strconv.Atoi(pageUser)
	if err != nil {
		page = 1
	}
	itemsPerPage, err := strconv.Atoi(itemsUser)
	if err != nil {
		itemsPerPage = 10
	}

	inicio := (page - 1) * itemsPerPage

	var resultado []Item
	if inicio >= 0 && inicio < len(items) {
		final := inicio + itemsPerPage
		if final > len(items) {
			final = len(items)
		}
		resultado = items[inicio:final]
	}

	json.NewEncoder(w).Encode(resultado)
}

func getItemId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parametros := mux.Vars(r)
	for _, item := range items {
		if item.ID == parametros["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode("Item no encontrado")
}

func getItemName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parametros := mux.Vars(r)
	for _, item := range items {
		if strings.EqualFold(item.Name, parametros["name"]) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode("Item no encontrado")
}
func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newItem Item
	rqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserte un item válido")
		return
	}
	json.Unmarshal(rqBody, &newItem)
	id := len(items) + 1
	newItem.ID = "item" + strconv.Itoa(id)
	items = append(items, newItem)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
	fmt.Println("Item creado exitosamente")
}
func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var UpdateItem Item
	parametros := mux.Vars(r)
	rqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserte un item válido")
		return
	}

	json.Unmarshal(rqBody, &UpdateItem)

	for i, item := range items {
		if item.ID == parametros["id"] {
			items = append(items[:i], items[i+1:]...)
			UpdateItem.ID = parametros["id"]
			items = append(items, UpdateItem)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(UpdateItem)
	fmt.Println("Item actualizado")
}
func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parametros := mux.Vars(r)
	for i, item := range items {
		if item.ID == parametros["id"] {
			items = append(items[:i], items[i+1:]...)
			json.NewEncoder(w).Encode("El Item fue eliminado exitosamente")
			return
		}
	}
	_, err := strconv.Atoi(parametros["id"])
	if err != nil {
		fmt.Println("ID inválido")
		return
	}
}

var items []Item

func main() {
	for i := 1; i <= 10; i++ {
		items = append(items, Item{ID: fmt.Sprintf("%d", i), Name: fmt.Sprintf("Item %d", i)})
	}

	router := mux.NewRouter()
	router.HandleFunc("/", raiz).Methods("GET")
	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items/{id}", getItemId).Methods("GET")
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
