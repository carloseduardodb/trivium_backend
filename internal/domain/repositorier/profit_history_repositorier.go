package repositorier

import "trivium/internal/domain/entity"

type ProfitHistoryRepository interface {
	Save(crypto entity.ProfitHistory) (entity.ProfitHistory, error)
	FindBySymbol(symbol string) ([]entity.ProfitHistory, error)
	FindAll() ([]entity.ProfitHistory, error)
}
