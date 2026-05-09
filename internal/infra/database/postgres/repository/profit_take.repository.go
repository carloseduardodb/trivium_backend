package repository

import (
	"fmt"
	"time"

	"trivium/internal/domain/entity"
	"trivium/internal/domain/repositorier"
	"trivium/internal/infra/database/postgres/connection"
)

type ProfitTakeRepository struct{}

func NewProfitTakeRepository() repositorier.ProfitTakeRepositorier {
	return &ProfitTakeRepository{}
}

func (r *ProfitTakeRepository) Save(profitTake entity.ProfitTake) (entity.ProfitTake, error) {
	db := connection.GetDB()

	query := `INSERT INTO profit_takes (position, amount_withdrawn, price_at_withdraw, remaining_value, withdraw_date, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := db.QueryRow(query,
		profitTake.Position,
		profitTake.AmountWithdrawn,
		profitTake.PriceAtWithdraw,
		profitTake.RemainingValue,
		profitTake.WithdrawDate,
		time.Now(),
		time.Now(),
	).Scan(&profitTake.ID)
	if err != nil {
		return entity.ProfitTake{}, err
	}

	return profitTake, nil
}

func (r *ProfitTakeRepository) FindById(id int64) (entity.ProfitTake, error) {
	db := connection.GetDB()

	var profitTake entity.ProfitTake
	err := db.QueryRowx(`
		SELECT id, position, amount_withdrawn, price_at_withdraw, remaining_value, 
		       withdraw_date, created_at, updated_at
		FROM profit_takes
		WHERE id = $1`, id).StructScan(&profitTake)
	if err != nil {
		return entity.ProfitTake{}, err
	}

	return profitTake, nil
}

func (r *ProfitTakeRepository) FindAll() ([]entity.ProfitTake, error) {
	db := connection.GetDB()

	var profitTakes []entity.ProfitTake
	err := db.Select(&profitTakes, `
		SELECT id, position, amount_withdrawn, price_at_withdraw, remaining_value, 
		       withdraw_date, created_at, updated_at
		FROM profit_takes
	`)
	if err != nil {
		return nil, err
	}

	return profitTakes, nil
}

func (r *ProfitTakeRepository) FindByPositionId(positionId int64) ([]entity.ProfitTake, error) {
	db := connection.GetDB()

	var profitTakes []entity.ProfitTake
	err := db.Select(&profitTakes, `
		SELECT id, position, amount_withdrawn, price_at_withdraw, remaining_value, 
		       withdraw_date, created_at, updated_at
		FROM profit_takes WHERE position = $1`, positionId)
	if err != nil {
		return nil, err
	}

	return profitTakes, nil
}

func (r *ProfitTakeRepository) Update(profitTake entity.ProfitTake) (entity.ProfitTake, error) {
	db := connection.GetDB()

	query := `
		UPDATE profit_takes 
		SET position = $1, amount_withdrawn = $2, price_at_withdraw = $3, 
		    remaining_value = $4, withdraw_date = $5, updated_at = $6
		WHERE id = $7
		RETURNING id, position, amount_withdrawn, price_at_withdraw, remaining_value, 
		          withdraw_date, created_at, updated_at`

	err := db.QueryRowx(query,
		profitTake.Position,
		profitTake.AmountWithdrawn,
		profitTake.PriceAtWithdraw,
		profitTake.RemainingValue,
		profitTake.WithdrawDate,
		time.Now(),
		profitTake.ID,
	).StructScan(&profitTake)

	if err != nil {
		return entity.ProfitTake{}, err
	}

	return profitTake, nil
}

func (r *ProfitTakeRepository) Delete(id int64) error {
	db := connection.GetDB()

	result, err := db.Exec("DELETE FROM profit_takes WHERE id = $1", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("profit take with id %d not found", id)
	}

	return nil
}
