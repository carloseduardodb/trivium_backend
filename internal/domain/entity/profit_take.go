package entity

import "time"

type ProfitTake struct {
	ID              int64     `json:"id"`
	Position        int64     `json:"position"`
	AmountWithdrawn float64   `json:"amount_withdrawn"`
	PriceAtWithdraw float64   `json:"price_at_withdraw"`
	RemainingValue  float64   `json:"remaining_value"`
	WithdrawDate    time.Time `json:"withdraw_date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
