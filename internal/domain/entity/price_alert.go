package entity

import "time"

type PriceAlert struct {
	ID             int64     `json:"id" db:"id"`
	UserID         int64     `json:"user_id" db:"user_id"`
	CryptoCurrency int64     `json:"crypto_currency" db:"crypto_currency"`
	Symbol         string    `json:"symbol" db:"symbol"`
	TargetPrice    float64   `json:"target_price" db:"target_price"`
	Direction      string    `json:"direction" db:"direction"` // "above" or "below"
	Active         bool      `json:"active" db:"active"`
	TriggeredAt    *time.Time `json:"triggered_at" db:"triggered_at"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}
