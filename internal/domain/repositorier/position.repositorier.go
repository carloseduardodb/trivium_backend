package repositorier

import (
	"trivium/internal/domain/entity"
)

type PositionRepositorier interface {
	Save(position entity.Position) (entity.Position, error)
	FindById(id int64) (entity.Position, error)
	FindAll() ([]entity.Position, error)
	Update(position entity.Position) (entity.Position, error)
	Delete(id int64) error
}
