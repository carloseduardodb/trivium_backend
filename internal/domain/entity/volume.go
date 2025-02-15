package entity

import "time"

type Volume struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Symbol    string    `json:"symbol"`
	CreatedAt time.Time `json:"created_at"`
}
