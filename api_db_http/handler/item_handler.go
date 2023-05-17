package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/ThaliaAC/labora-api/api_db_http/model"
	"github.com/ThaliaAC/labora-api/api_db_http/service"
	"github.com/gorilla/mux"
)

var items []model.Item

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items, _ := service.GetItemsDb()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

func GetItemsByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	idInput, idErr := strconv.Atoi(params["id"])
	if idErr != nil {
		log.Fatal(idErr)
	}

	item := service.GetItemsById(idInput)
	if item == nil {
		json.NewEncoder(w).Encode("Item not found")
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func GetItemsByNameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	item := service.GetItemsByName(params["customerName"])
	if item == nil {
		json.NewEncoder(w).Encode("Item not found")
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newItem model.Item
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Customer")
		return
	}

	json.Unmarshal(reqBody, &newItem)

	err = service.CreateItem(newItem)
	if err != nil {
		fmt.Fprintf(w, "Enter Valid Data")
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
	fmt.Println("Customer successfully created")
}

func UpdateItemHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idInput, idErr := strconv.Atoi(params["id"])
	if idErr != nil {
		log.Fatal(idErr)
	}
	var updatedItem model.Item
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Customer")
		return
	}

	json.Unmarshal(reqBody, &updatedItem)

	finalOutput, errOutput := service.UpdateItem(idInput, updatedItem)
	if errOutput != nil {
		fmt.Fprintf(w, "Enter Valid Data")
		return
	}
	for index, item := range items {
		if item.ID == params["id"] {
			items = append(items[:index], items[index+1:]...)
			updatedItem.ID = params["id"]
			items = append(items, updatedItem)
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(finalOutput)
		fmt.Println("Customer successfully created")
	}
}

func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	idInput, idErr := strconv.Atoi(params["id"])
	if idErr != nil {
		log.Fatal(idErr)
	}
	err := service.DeleteItem(idInput)
	if err != nil {
		fmt.Fprintf(w, "Enter Valid Data")
		return
	}
	for index, item := range items {
		if item.ID == params["id"] {
			items = append(items[:index], items[index+1:]...)
			json.NewEncoder(w).Encode("The Customer has been remove succesfully")
			return
		}
	}
}
