package usecase

import "trivium/internal/domain/entity"

type AddCryptoCurrencyUseCase struct{}

func NewAddCryptoCurrencyUseCase() *AddCryptoCurrencyUseCase {
	return &AddCryptoCurrencyUseCase{}
}

func (a *AddCryptoCurrencyUseCase) AddCryptoCurrency(cryptoCurrency *entity.CryptoCurrency) error {
	return nil
}
