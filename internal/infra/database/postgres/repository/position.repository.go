package repository

import (
	"fmt"
	"time"

	"trivium/internal/domain/entity"
	"trivium/internal/domain/repositorier"
	"trivium/internal/infra/database/postgres/connection"
)

type PositionRepository struct{}

func NewPositionRepository() repositorier.PositionRepositorier {
	return &PositionRepository{}
}

func (r *PositionRepository) Save(position entity.Position) (entity.Position, error) {
	db := connection.GetDB()

	query := `INSERT INTO positions (
		user_id, crypto_currency, quantity, purchase_price, invested_amount, 
		purchase_date, last_profit_price, status, created_at, updated_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`

	err := db.QueryRow(query,
		position.UserID,
		position.CryptoCurrency,
		position.Quantity,
		position.PurchasePrice,
		position.InvestedAmount,
		position.PurchaseDate,
		position.LastProfitPrice,
		position.Status,
		time.Now(),
		time.Now(),
	).Scan(&position.ID)
	if err != nil {
		return entity.Position{}, err
	}

	return position, nil
}

func (r *PositionRepository) FindById(id int64) (entity.Position, error) {
	db := connection.GetDB()

	var position entity.Position
	err := db.QueryRowx(`
		SELECT id, user_id, crypto_currency, quantity, purchase_price, invested_amount,
		       purchase_date, last_profit_price, status, created_at, updated_at
		FROM positions
		WHERE id = $1`, id).StructScan(&position)
	if err != nil {
		return entity.Position{}, err
	}

	return position, nil
}

func (r *PositionRepository) FindAll() ([]entity.Position, error) {
	db := connection.GetDB()

	var positions []entity.Position
	err := db.Select(&positions, `
		SELECT id, user_id, crypto_currency, quantity, purchase_price, invested_amount,
		       purchase_date, last_profit_price, status, created_at, updated_at
		FROM positions`)
	if err != nil {
		return nil, err
	}

	return positions, nil
}

func (r *PositionRepository) FindByUserId(userId int64) ([]entity.Position, error) {
	db := connection.GetDB()

	var positions []entity.Position
	err := db.Select(&positions, `
		SELECT id, user_id, crypto_currency, quantity, purchase_price, invested_amount,
		       purchase_date, last_profit_price, status, created_at, updated_at
		FROM positions WHERE user_id = $1`, userId)
	if err != nil {
		return nil, err
	}

	return positions, nil
}

func (r *PositionRepository) Update(position entity.Position) (entity.Position, error) {
	db := connection.GetDB()

	query := `
		UPDATE positions 
		SET crypto_currency = $1, quantity = $2, purchase_price = $3,
		    invested_amount = $4, purchase_date = $5, last_profit_price = $6,
		    status = $7, updated_at = $8
		WHERE id = $9
		RETURNING id, user_id, crypto_currency, quantity, purchase_price, invested_amount,
		          purchase_date, last_profit_price, status, created_at, updated_at`

	err := db.QueryRowx(query,
		position.CryptoCurrency,
		position.Quantity,
		position.PurchasePrice,
		position.InvestedAmount,
		position.PurchaseDate,
		position.LastProfitPrice,
		position.Status,
		time.Now(),
		position.ID,
	).StructScan(&position)

	if err != nil {
		return entity.Position{}, err
	}

	return position, nil
}

func (r *PositionRepository) Delete(id int64) error {
	db := connection.GetDB()

	result, err := db.Exec("DELETE FROM positions WHERE id = $1", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("position with id %d not found", id)
	}

	return nil
}
