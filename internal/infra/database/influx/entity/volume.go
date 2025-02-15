package entity

import "time"

type Volume struct {
	Id        int64     `db:"id"`
	Name      string    `db:"name"`
	Price     float64   `db:"price"`
	Symbol    string    `db:"symbol"`
	CreatedAt time.Time `db:"created_at"`
}
