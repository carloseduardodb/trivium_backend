package dto

type CreateCryptoCurrency struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type UpdateCryptoCurrency struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}
