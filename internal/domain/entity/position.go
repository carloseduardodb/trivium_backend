package entity

import "time"

type Position struct {
	ID              int64     `json:"id" db:"id"`
	UserID          int64     `json:"user_id" db:"user_id"`
	CryptoCurrency  int64     `json:"crypto_currency" db:"crypto_currency"`
	Quantity        float64   `json:"quantity" db:"quantity"`
	PurchasePrice   float64   `json:"purchase_price" db:"purchase_price"`
	InvestedAmount  float64   `json:"invested_amount" db:"invested_amount"`
	PurchaseDate    time.Time `json:"purchase_date" db:"purchase_date"`
	LastProfitPrice float64   `json:"last_profit_price" db:"last_profit_price"`
	Status          string    `json:"status" db:"status"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}
