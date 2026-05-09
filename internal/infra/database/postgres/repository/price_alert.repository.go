package repository

import (
	"fmt"
	"time"

	"trivium/internal/domain/entity"
	"trivium/internal/domain/repositorier"
	"trivium/internal/infra/database/postgres/connection"
)

type PriceAlertRepository struct{}

func NewPriceAlertRepository() repositorier.PriceAlertRepositorier {
	return &PriceAlertRepository{}
}

func (r *PriceAlertRepository) Save(alert entity.PriceAlert) (entity.PriceAlert, error) {
	db := connection.GetDB()

	query := `INSERT INTO price_alerts (user_id, crypto_currency, symbol, target_price, direction, active, created_at)
			  VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := db.QueryRow(query,
		alert.UserID,
		alert.CryptoCurrency,
		alert.Symbol,
		alert.TargetPrice,
		alert.Direction,
		true,
		time.Now(),
	).Scan(&alert.ID)
	if err != nil {
		return entity.PriceAlert{}, err
	}

	alert.Active = true
	return alert, nil
}

func (r *PriceAlertRepository) FindByUserId(userId int64) ([]entity.PriceAlert, error) {
	db := connection.GetDB()

	var alerts []entity.PriceAlert
	err := db.Select(&alerts, `
		SELECT id, user_id, crypto_currency, symbol, target_price, direction, active, triggered_at, created_at
		FROM price_alerts WHERE user_id = $1 ORDER BY created_at DESC`, userId)
	if err != nil {
		return nil, err
	}

	return alerts, nil
}

func (r *PriceAlertRepository) FindActive() ([]entity.PriceAlert, error) {
	db := connection.GetDB()

	var alerts []entity.PriceAlert
	err := db.Select(&alerts, `
		SELECT id, user_id, crypto_currency, symbol, target_price, direction, active, triggered_at, created_at
		FROM price_alerts WHERE active = true`)
	if err != nil {
		return nil, err
	}

	return alerts, nil
}

func (r *PriceAlertRepository) Deactivate(id int64) error {
	db := connection.GetDB()

	now := time.Now()
	result, err := db.Exec(`UPDATE price_alerts SET active = false, triggered_at = $1 WHERE id = $2`, now, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("price alert with id %d not found", id)
	}

	return nil
}

func (r *PriceAlertRepository) Delete(id int64) error {
	db := connection.GetDB()

	result, err := db.Exec("DELETE FROM price_alerts WHERE id = $1", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("price alert with id %d not found", id)
	}

	return nil
}
