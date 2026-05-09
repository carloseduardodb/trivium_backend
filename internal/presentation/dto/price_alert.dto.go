package dto

type CreatePriceAlert struct {
	CryptoCurrency int64   `json:"crypto_currency"`
	TargetPrice    float64 `json:"target_price"`
	Direction      string  `json:"direction"` // "above" or "below"
}
