package model

import (
	"time"
)

type Item struct {
	ID           string    `json:"id"`
	CustomerName string    `json:"customerName"`
	OrderDate    time.Time `json:"orderDate"`
	Product      string    `json:"product"`
	Quantity     int       `json:"quantity"`
	Price        int       `json:"price"`
	Details      string    `json:"details,omitempty"`
}

func (item Item) TotalPrice() int {
	totalPrice := item.Price * item.Quantity
	return totalPrice
}
