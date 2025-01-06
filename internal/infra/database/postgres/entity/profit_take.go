package entity

import "time"

type ProfitTake struct {
	ID              int64     `db:"id"`
	Position        int64     `db:"position"`
	AmountWithdrawn float64   `db:"amount_withdrawn"`
	PriceAtWithdraw float64   `db:"price_at_withdraw"`
	RemainingValue  float64   `db:"remaining_value"`
	WithdrawDate    time.Time `db:"withdraw_date"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}
