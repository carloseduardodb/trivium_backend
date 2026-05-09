package dto

import "time"

type CreatePosition struct {
	CryptoCurrency int64     `json:"crypto_currency"`
	Quantity       float64   `json:"quantity"`
	PurchasePrice  float64   `json:"purchase_price"`
	PurchaseDate   time.Time `json:"purchase_date"`
}

type UpdatePosition struct {
	ID              int64   `json:"id"`
	Quantity        float64 `json:"quantity"`
	PurchasePrice   float64 `json:"purchase_price"`
	LastProfitPrice float64 `json:"last_profit_price"`
	Status          string  `json:"status"`
}
