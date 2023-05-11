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

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items, _ := service.GetUsersDb()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

func GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	idInput, idErr := strconv.Atoi(params["id"])
	if idErr != nil {
		log.Fatal(idErr)
	}

	item := service.GetUserById(idInput)
	if item == nil {
		json.NewEncoder(w).Encode("Item not found")
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func GetUserByNameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	item := service.GetUserByName(params["customerName"])
	if item == nil {
		json.NewEncoder(w).Encode("Item not found")
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

/*func CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newItem model.Item
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Customer")
		return
	}

	json.Unmarshal(reqBody, &newItem)

	id := len(items) + 1
	newItem.ID = strconv.Itoa(id)

	items = append(items, newItem)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
	fmt.Println("Customer successfully created")
}*/

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updatedItem model.Item
	params := mux.Vars(r)
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter Valid Data")
		return
	}

	json.Unmarshal(reqBody, &updatedItem)

	for index, item := range items {
		if item.ID == params["id"] {
			items = append(items[:index], items[index+1:]...)
			updatedItem.ID = params["id"]
			items = append(items, updatedItem)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedItem)
	fmt.Println("Customer has been updated successfully")
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range items {
		if item.ID == params["id"] {
			items = append(items[:index], items[index+1:]...)
			json.NewEncoder(w).Encode("The Customer has been remove succesfully")
			return
		}
	}
	_, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}
}

func GetUserDetailsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	idInput, idErr := strconv.Atoi(params["id"])
	if idErr != nil {
		log.Fatal(idErr)
	}

	item := service.GetUserDetails(idInput)
	if item == nil {
		json.NewEncoder(w).Encode("Item not found")
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}
