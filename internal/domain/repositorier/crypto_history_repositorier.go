package repositorier

import "trivium/internal/domain/entity"

type CryptoHistoryRepository interface {
	Save(crypto entity.CryptoHistory) (entity.CryptoHistory, error)
	FindBySymbol(symbol string) ([]entity.CryptoHistory, error)
	FindAll() ([]entity.CryptoHistory, error)
}
