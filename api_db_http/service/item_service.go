package service

import (
	"log"

	"github.com/ThaliaAC/labora-api/api_db_http/db"
	"github.com/ThaliaAC/labora-api/api_db_http/model"
)

func GetUsersDb() ([]model.Item, error) {
	items := make([]model.Item, 0)
	rows, err := db.Db.Query("SELECT * FROM items")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var item model.Item
		err := rows.Scan(&item.ID, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price, &item.Details)
		items = append(items, item)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return items, nil
}

func GetUserById(id int) *model.Item {
	row := db.Db.QueryRow("SELECT id, customer_name, order_date, product, quantity, price FROM items WHERE id = $1", id)
	var item model.Item
	err := row.Scan(&item.ID, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price)
	if err != nil {
		log.Fatal(err)
	}
	err = row.Err()
	if err != nil {
		log.Fatal(err)
	}
	return &item
}

func GetUserByName(customerName string) *model.Item {
	row := db.Db.QueryRow("SELECT id, customer_name, order_date, product, quantity, price FROM items WHERE customerName = $1", customerName)
	var item model.Item
	err := row.Scan(&item.ID, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price)
	if err != nil {
		log.Fatal(err)
	}
	err = row.Err()
	if err != nil {
		log.Fatal(err)
	}
	return &item
}

/*func CreateItem() {

}*/

func UpdateUser(id int, item model.Item) (model.Item, error) {
	var updatedItem model.Item
	row := db.Db.QueryRow("UPDATE items SET customer name $1, order_date = $2, product = $3, quantity = $4, price = $5, details = $6 WHERE id = $7 RETURNING *",
		item.CustomerName, item.OrderDate, item.Product, item.Quantity, item.Price, item.Details, id)
	err := row.Scan(&updatedItem.ID, &updatedItem.CustomerName, &updatedItem.OrderDate, &updatedItem.Product, &updatedItem.Quantity, &updatedItem.Price, &updatedItem.Details)
	if err != nil {
		log.Fatal(err)
		return updatedItem, err
	}
	return updatedItem, nil
}

func DeleteUser(id int) error {
	_, err := db.Db.Exec("DELETE FROM items WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func GetUserDetails(id int) *model.Item {
	row := db.Db.QueryRow("SELECT id, customer_name, product, quantity, price, details FROM items WHERE id = $1", id)
	var item model.Item
	err := row.Scan(&item.ID, &item.CustomerName, &item.Product, &item.Quantity, &item.Price, &item.Details)
	if err != nil {
		log.Fatal(err)
	}
	err = row.Err()
	if err != nil {
		log.Fatal(err)
	}
	return &item
}
