package entity

import "time"

type ProfitHistory struct {
	Id          int64     `json:"id"`
	Symbol      string    `json:"symbol"`
	CryptoPrice float64   `json:"crypto_price"`
	UserId      int64     `json:"user_id"`
	Profit      float64   `json:"profit"`
	CreatedAt   time.Time `json:"created_at"`
}
