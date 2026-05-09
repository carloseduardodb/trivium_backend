package usecase

import (
	"fmt"

	"trivium/internal/domain/entity"
	"trivium/internal/domain/repositorier"
	"trivium/internal/presentation/dto"
)

type ProfitTakeUseCase struct {
	profitTakeRepo repositorier.ProfitTakeRepositorier
	positionRepo   repositorier.PositionRepositorier
}

func NewProfitTakeUseCase(
	profitTakeRepo repositorier.ProfitTakeRepositorier,
	positionRepo repositorier.PositionRepositorier,
) *ProfitTakeUseCase {
	return &ProfitTakeUseCase{
		profitTakeRepo: profitTakeRepo,
		positionRepo:   positionRepo,
	}
}

func (p *ProfitTakeUseCase) Create(userId int64, input *dto.CreateProfitTake) (*entity.ProfitTake, error) {
	if input.PositionID == 0 {
		return nil, fmt.Errorf("position_id is required")
	}
	if input.AmountWithdrawn <= 0 {
		return nil, fmt.Errorf("amount_withdrawn must be greater than zero")
	}
	if input.PriceAtWithdraw <= 0 {
		return nil, fmt.Errorf("price_at_withdraw must be greater than zero")
	}

	position, err := p.positionRepo.FindById(input.PositionID)
	if err != nil {
		return nil, fmt.Errorf("position not found: %w", err)
	}

	if position.UserID != userId {
		return nil, fmt.Errorf("unauthorized: position does not belong to user")
	}

	if position.Status == "closed" {
		return nil, fmt.Errorf("cannot take profit from a closed position")
	}

	currentValue := position.Quantity * input.PriceAtWithdraw
	remainingValue := currentValue - input.AmountWithdrawn

	if remainingValue < 0 {
		return nil, fmt.Errorf("amount_withdrawn exceeds position value")
	}

	profitTake := entity.ProfitTake{
		Position:        input.PositionID,
		AmountWithdrawn: input.AmountWithdrawn,
		PriceAtWithdraw: input.PriceAtWithdraw,
		RemainingValue:  remainingValue,
		WithdrawDate:    input.WithdrawDate,
	}

	saved, err := p.profitTakeRepo.Save(profitTake)
	if err != nil {
		return nil, fmt.Errorf("error saving profit take: %w", err)
	}

	position.LastProfitPrice = input.PriceAtWithdraw
	_, err = p.positionRepo.Update(position)
	if err != nil {
		return nil, fmt.Errorf("error updating position: %w", err)
	}

	return &saved, nil
}

func (p *ProfitTakeUseCase) FindByPositionId(positionId int64) ([]entity.ProfitTake, error) {
	return p.profitTakeRepo.FindByPositionId(positionId)
}

func (p *ProfitTakeUseCase) Delete(id int64, userId int64) error {
	profitTake, err := p.profitTakeRepo.FindById(id)
	if err != nil {
		return fmt.Errorf("profit take not found: %w", err)
	}

	position, err := p.positionRepo.FindById(profitTake.Position)
	if err != nil {
		return fmt.Errorf("position not found: %w", err)
	}

	if position.UserID != userId {
		return fmt.Errorf("unauthorized: position does not belong to user")
	}

	return p.profitTakeRepo.Delete(id)
}
