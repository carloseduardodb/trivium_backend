package repositorier

import "trivium/internal/domain/entity"

type CryptoCurrencyRepositorier interface {
	Save(cryptoCurrency entity.CryptoCurrency) (entity.CryptoCurrency, error)
	FindById(id int64) (entity.CryptoCurrency, error)
	FindAll() ([]entity.CryptoCurrency, error)
	Update(cryptoCurrency entity.CryptoCurrency) (entity.CryptoCurrency, error)
	Delete(id int64) error
}
