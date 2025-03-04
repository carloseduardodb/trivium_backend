package entity

import "time"

type Volume struct {
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Symbol    string    `json:"symbol"`
	CreatedAt time.Time `json:"created_at"`
}
