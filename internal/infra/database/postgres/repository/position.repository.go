package repository

import (
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
	defer connection.CloseDB()

	query := `INSERT INTO positions (
		crypto_currency, quantity, purchase_price, invested_amount, 
		purchase_date, last_profit_price, status, user, created_at, updated_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`

	err := db.QueryRow(query,
		position.CryptoCurrency,
		position.Quantity,
		position.PurchasePrice,
		position.InvestedAmount,
		position.PurchaseDate,
		position.LastProfitPrice,
		position.Status,
		position.User,
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
	defer connection.CloseDB()

	var position entity.Position
	err := db.QueryRowx(`
		SELECT id, crypto_currency, quantity, purchase_price, invested_amount,
		       purchase_date, last_profit_price, status, user, created_at, updated_at
		FROM positions
		WHERE id = $1`, id).StructScan(&position)
	if err != nil {
		return entity.Position{}, err
	}

	return position, nil
}

func (r *PositionRepository) FindAll() ([]entity.Position, error) {
	db := connection.GetDB()
	defer connection.CloseDB()

	var positions []entity.Position
	err := db.Select(&positions, `
		SELECT id, crypto_currency, quantity, purchase_price, invested_amount,
		       purchase_date, last_profit_price, status, user, created_at, updated_at
		FROM positions`)
	if err != nil {
		return nil, err
	}

	return positions, nil
}

func (r *PositionRepository) Update(position entity.Position) (entity.Position, error) {
	db := connection.GetDB()
	defer connection.CloseDB()

	query := `
		UPDATE positions 
		SET crypto_currency = $1, quantity = $2, purchase_price = $3,
		    invested_amount = $4, purchase_date = $5, last_profit_price = $6,
		    status = $7, user = $8, updated_at = $9
		WHERE id = $10
		RETURNING id, crypto_currency, quantity, purchase_price, invested_amount,
		          purchase_date, last_profit_price, status, user, created_at, updated_at`

	err := db.QueryRowx(query,
		position.CryptoCurrency,
		position.Quantity,
		position.PurchasePrice,
		position.InvestedAmount,
		position.PurchaseDate,
		position.LastProfitPrice,
		position.Status,
		position.User,
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
	defer connection.CloseDB()

	_, err := db.Exec("DELETE FROM positions WHERE id = $1", id)
	return err
}
