package entity

import "time"

type CryptoHistory struct {
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Symbol    string    `json:"symbol"`
	CreatedAt time.Time `json:"created_at"`
}
