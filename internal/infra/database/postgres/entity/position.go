package entity

import "time"

type Position struct {
	ID              int64     `db:"id"`
	CryptoCurrency  int64     `db:"crypto_currency"`
	Quantity        float64   `db:"quantity"`
	PurchasePrice   float64   `db:"purchase_price"`
	InvestedAmount  float64   `db:"invested_amount"`
	PurchaseDate    time.Time `db:"purchase_date"`
	LastProfitPrice float64   `db:"last_profit_price"`
	Status          string    `db:"status"`
	User            User      `db:"user"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}
