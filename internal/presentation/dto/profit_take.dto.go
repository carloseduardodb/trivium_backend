package dto

import "time"

type CreateProfitTake struct {
	PositionID      int64     `json:"position_id"`
	AmountWithdrawn float64   `json:"amount_withdrawn"`
	PriceAtWithdraw float64   `json:"price_at_withdraw"`
	WithdrawDate    time.Time `json:"withdraw_date"`
}
