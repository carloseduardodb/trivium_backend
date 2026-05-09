package usecase

import (
	"fmt"

	"trivium/internal/domain/entity"
	"trivium/internal/domain/repositorier"
	"trivium/internal/presentation/dto"
)

type CryptoCurrencyUseCase struct {
	cryptoRepo repositorier.CryptoCurrencyRepositorier
}

func NewCryptoCurrencyUseCase(cryptoRepo repositorier.CryptoCurrencyRepositorier) *CryptoCurrencyUseCase {
	return &CryptoCurrencyUseCase{
		cryptoRepo: cryptoRepo,
	}
}

func (c *CryptoCurrencyUseCase) Create(input *dto.CreateCryptoCurrency) (*entity.CryptoCurrency, error) {
	if input.Name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if input.Symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	crypto := entity.CryptoCurrency{
		Name:   input.Name,
		Symbol: input.Symbol,
	}

	saved, err := c.cryptoRepo.Save(crypto)
	if err != nil {
		return nil, fmt.Errorf("error saving cryptocurrency: %w", err)
	}

	return &saved, nil
}

func (c *CryptoCurrencyUseCase) FindAll() ([]entity.CryptoCurrency, error) {
	return c.cryptoRepo.FindAll()
}

func (c *CryptoCurrencyUseCase) FindById(id int64) (*entity.CryptoCurrency, error) {
	crypto, err := c.cryptoRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("cryptocurrency not found: %w", err)
	}
	return &crypto, nil
}

func (c *CryptoCurrencyUseCase) Update(input *dto.UpdateCryptoCurrency) (*entity.CryptoCurrency, error) {
	if input.ID == 0 {
		return nil, fmt.Errorf("id is required")
	}

	crypto := entity.CryptoCurrency{
		ID:     input.ID,
		Name:   input.Name,
		Symbol: input.Symbol,
	}

	updated, err := c.cryptoRepo.Update(crypto)
	if err != nil {
		return nil, fmt.Errorf("error updating cryptocurrency: %w", err)
	}

	return &updated, nil
}

func (c *CryptoCurrencyUseCase) Delete(id int64) error {
	return c.cryptoRepo.Delete(id)
}
