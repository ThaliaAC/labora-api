package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ThaliaAC/labora-api/api_db_http/db"
	"github.com/ThaliaAC/labora-api/api_db_http/model"
)

var Items []model.Item

// Función para obtener todos los elementos
func GetUsersDB() []model.Item {
	rows, err := db.DB.Query("SELECT * FROM items")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var item model.Item
		err := rows.Scan(&item.ID, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price, &item.Details)
		if err != nil {
			log.Fatal(err)
		}
		Items = append(Items, item)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return Items
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items := GetUsersDB()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

func GetUserIdDb() []model.Item {
	rows, err := db.DB.Query("SELECT ID FROM items")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var item model.Item
		err := rows.Scan(&item.ID)
		if err != nil {
			log.Fatal(err)
		}
		Items = append(Items, item)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return Items
}

// Función para obtener un elemento específico
func GetUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := GetUserIdDb()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}

/*
// Función para buscar un elemento por nombre
func GetUserName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range Items {
		if strings.EqualFold(item.Name, params["name"]) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode("Customer not found")
}

// Función para crear un nuevo elemento
func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Recibo la información
	var newItem model.Item
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Customer")
		return
	}
	//Asigno esa información a la variable newItem
	json.Unmarshal(reqBody, &newItem)

	//Automatizo incremento de ID
	id := len(Items) + 1
	newItem.ID = strconv.Itoa(id)
	//Guardo newItem en la lista de items
	Items = append(Items, newItem)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
	fmt.Println("Customer successfully created")
}

// Función para actualizar un elemento existente
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updatedItem model.Item
	params := mux.Vars(r)
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter Valid Data")
		return
	}

	json.Unmarshal(reqBody, &updatedItem)

	for index, item := range Items {
		if item.ID == params["id"] {
			Items = append(Items[:index], Items[index+1:]...)
			updatedItem.ID = params["id"]
			Items = append(Items, updatedItem)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedItem)
	fmt.Println("Customer has been updated successfully")
}

// Función para eliminar un elemento
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Items {
		if item.ID == params["id"] {
			Items = append(Items[:index], Items[index+1:]...)
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

func GetUserDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range Items {
		if strings.EqualFold(item.Details, params["details"]) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode("Customer not found")
}*/
