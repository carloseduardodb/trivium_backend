package entity

import "time"

type Position struct {
	ID              int64     `json:"id"`
	CryptoCurrency  int64     `json:"crypto_currency"`
	Quantity        float64   `json:"quantity"`
	PurchasePrice   float64   `json:"purchase_price"`
	InvestedAmount  float64   `json:"invested_amount"`
	PurchaseDate    time.Time `json:"purchase_date"`
	LastProfitPrice float64   `json:"last_profit_price"`
	Status          string    `json:"status"`
	User            User      `json:"user"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
