package entity

import "time"

type CryptoCurrency struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Symbol    string    `db:"symbol"`
	User      User      `db:"user"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
