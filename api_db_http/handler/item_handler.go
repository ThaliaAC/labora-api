package handler

import (
	"encoding/json"
	"fmt"
	"io"
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
	items, err := service.GetItemsDb()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

func GetItemsByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	idInput, idErr := strconv.Atoi(params["id"])
	if idErr != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(idErr)
		return
	}

	item, _ := service.GetItemsById(idInput)

	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(item)
}

func GetItemsByNameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	item, nameErr := service.GetItemsByName(params["customerName"])
	if nameErr != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(nameErr)
		return
	}
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(item)
}

func CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newItem model.Item
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	json.Unmarshal(reqBody, &newItem)

	_ = service.CreateItem(newItem)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
	fmt.Println("Customer successfully created")
}

func UpdateItemHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idInput, idErr := strconv.Atoi(params["id"])
	if idErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(idErr)
		return
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
		w.WriteHeader(http.StatusNotModified)
		json.NewEncoder(w).Encode(errOutput)
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
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(idErr)
		return
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
