package repositorier

import (
	"trivium/internal/domain/entity"
)

type ProfitTakeRepositorier interface {
	Save(profitTake entity.ProfitTake) (entity.ProfitTake, error)
	FindById(id int64) (entity.ProfitTake, error)
	FindAll() ([]entity.ProfitTake, error)
	Update(profitTake entity.ProfitTake) (entity.ProfitTake, error)
	Delete(id int64) error
}
