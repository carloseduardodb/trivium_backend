package repository

import (
	"trivium/internal/domain/entity"
	"trivium/internal/domain/repositorier"
	"trivium/internal/infra/database/postgres/connection"
)

type CryptoCurrencyRepository struct{}

func NewCryptoCurrencyRepository() repositorier.CryptoCurrencyRepositorier {
	return &CryptoCurrencyRepository{}
}

func (r *CryptoCurrencyRepository) Save(cryptoCurrency entity.CryptoCurrency) (entity.CryptoCurrency, error) {
	db := connection.GetDB()
	defer connection.CloseDB()

	query := `INSERT INTO cryptocurrencies (name, symbol) VALUES ($1, $2) RETURNING id`
	err := db.QueryRow(query, cryptoCurrency.Name, cryptoCurrency.Symbol).Scan(&cryptoCurrency.ID)
	if err != nil {
		return entity.CryptoCurrency{}, err
	}

	return cryptoCurrency, nil
}

func (r *CryptoCurrencyRepository) FindById(id int64) (entity.CryptoCurrency, error) {
	db := connection.GetDB()
	defer connection.CloseDB()

	var cryptoCurrency entity.CryptoCurrency
	err := db.QueryRowx("SELECT id, name, symbol FROM cryptocurrencies WHERE id = $1", id).StructScan(&cryptoCurrency)
	if err != nil {
		return entity.CryptoCurrency{}, err
	}

	return cryptoCurrency, nil
}

func (r *CryptoCurrencyRepository) FindAll() ([]entity.CryptoCurrency, error) {
	db := connection.GetDB()
	defer connection.CloseDB()

	cryptoCurrencies := []entity.CryptoCurrency{}
	err := db.Select(&cryptoCurrencies, "SELECT id, name, symbol FROM cryptocurrencies")
	if err != nil {
		return nil, err
	}

	return cryptoCurrencies, nil
}

func (r *CryptoCurrencyRepository) Update(cryptoCurrency entity.CryptoCurrency) (entity.CryptoCurrency, error) {
	db := connection.GetDB()
	defer connection.CloseDB()

	query := `UPDATE cryptocurrencies SET name = $1, symbol = $2, WHERE id = $3`
	result, err := db.Exec(query, cryptoCurrency.Name, cryptoCurrency.Symbol, cryptoCurrency.ID)
	if err != nil {
		return entity.CryptoCurrency{}, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return entity.CryptoCurrency{}, err
	}
	if rows == 0 {
		return entity.CryptoCurrency{}, err
	}

	return cryptoCurrency, nil
}

func (r *CryptoCurrencyRepository) Delete(id int64) error {
	db := connection.GetDB()
	defer connection.CloseDB()

	result, err := db.Exec("DELETE FROM cryptocurrencies WHERE id = $1", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return err
	}

	return nil
}
