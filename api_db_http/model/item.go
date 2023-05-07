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
	Price        float32   `json:"price"`
	Details      string    `json:"details,omitempty"`
}
