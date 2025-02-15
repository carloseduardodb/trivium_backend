package repositorier

import "trivium/internal/domain/entity"

type VolumeRepository interface {
	Save(crypto entity.Volume) (entity.Volume, error)
	FindBySymbol(symbol string) ([]entity.Volume, error)
	FindAll() ([]entity.Volume, error)
}
