package entity

import "time"

type CryptoCurrency struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Symbol    string    `json:"symbol"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
