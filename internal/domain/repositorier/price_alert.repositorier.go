package repositorier

import "trivium/internal/domain/entity"

type PriceAlertRepositorier interface {
	Save(alert entity.PriceAlert) (entity.PriceAlert, error)
	FindByUserId(userId int64) ([]entity.PriceAlert, error)
	FindActive() ([]entity.PriceAlert, error)
	Deactivate(id int64) error
	Delete(id int64) error
}
