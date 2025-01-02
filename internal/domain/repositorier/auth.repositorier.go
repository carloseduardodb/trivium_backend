package repositorier

import "trivium/internal/domain/entity"

type AuthRepositorier interface {
	ValidateToken(token string) (bool, error)
	ConvertTokenInUser(token string) (*entity.User, error)
}
