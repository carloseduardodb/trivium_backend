package usecase

import (
	"fmt"

	"trivium/internal/domain/entity"
	"trivium/internal/domain/repositorier"
	"trivium/internal/presentation/dto"
)

type PositionUseCase struct {
	positionRepo repositorier.PositionRepositorier
	cryptoRepo   repositorier.CryptoCurrencyRepositorier
}

func NewPositionUseCase(
	positionRepo repositorier.PositionRepositorier,
	cryptoRepo repositorier.CryptoCurrencyRepositorier,
) *PositionUseCase {
	return &PositionUseCase{
		positionRepo: positionRepo,
		cryptoRepo:   cryptoRepo,
	}
}

func (p *PositionUseCase) Create(userId int64, input *dto.CreatePosition) (*entity.Position, error) {
	if input.CryptoCurrency == 0 {
		return nil, fmt.Errorf("crypto_currency is required")
	}
	if input.Quantity <= 0 {
		return nil, fmt.Errorf("quantity must be greater than zero")
	}
	if input.PurchasePrice <= 0 {
		return nil, fmt.Errorf("purchase_price must be greater than zero")
	}

	_, err := p.cryptoRepo.FindById(input.CryptoCurrency)
	if err != nil {
		return nil, fmt.Errorf("cryptocurrency not found: %w", err)
	}

	position := entity.Position{
		UserID:         userId,
		CryptoCurrency: input.CryptoCurrency,
		Quantity:       input.Quantity,
		PurchasePrice:  input.PurchasePrice,
		InvestedAmount: input.Quantity * input.PurchasePrice,
		PurchaseDate:   input.PurchaseDate,
		Status:         "active",
	}

	saved, err := p.positionRepo.Save(position)
	if err != nil {
		return nil, fmt.Errorf("error saving position: %w", err)
	}

	return &saved, nil
}

func (p *PositionUseCase) FindByUserId(userId int64) ([]entity.Position, error) {
	return p.positionRepo.FindByUserId(userId)
}

func (p *PositionUseCase) FindById(id int64) (*entity.Position, error) {
	position, err := p.positionRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("position not found: %w", err)
	}
	return &position, nil
}

func (p *PositionUseCase) Close(id int64, userId int64) (*entity.Position, error) {
	position, err := p.positionRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("position not found: %w", err)
	}

	if position.UserID != userId {
		return nil, fmt.Errorf("unauthorized: position does not belong to user")
	}

	if position.Status == "closed" {
		return nil, fmt.Errorf("position is already closed")
	}

	position.Status = "closed"
	updated, err := p.positionRepo.Update(position)
	if err != nil {
		return nil, fmt.Errorf("error closing position: %w", err)
	}

	return &updated, nil
}

func (p *PositionUseCase) Delete(id int64, userId int64) error {
	position, err := p.positionRepo.FindById(id)
	if err != nil {
		return fmt.Errorf("position not found: %w", err)
	}

	if position.UserID != userId {
		return fmt.Errorf("unauthorized: position does not belong to user")
	}

	return p.positionRepo.Delete(id)
}
