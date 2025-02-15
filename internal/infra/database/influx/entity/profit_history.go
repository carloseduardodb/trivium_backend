package entity

import "time"

type ProfitHistory struct {
	Id          int64     `db:"id"`
	Symbol      string    `db:"symbol"`
	CryptoPrice float64   `db:"crypto_price"`
	Profit      float64   `db:"profit"`
	UserId      int64     `db:"user_id"`
	CreatedAt   time.Time `db:"created_at"`
}
