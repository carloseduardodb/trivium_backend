package repositorier

import "trivium/internal/domain/entity"

type UserRepositorier interface {
	Save(user entity.User) (entity.User, error)
	FindById(id int64) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	FindAll() ([]entity.User, error)
	Update(user entity.User) (entity.User, error)
	Delete(id int64) error
}
