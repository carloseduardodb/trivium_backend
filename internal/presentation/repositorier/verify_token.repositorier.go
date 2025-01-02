package presentation_repositorier

import (
	"context"
	"trivium/internal/presentation/dto"
)

type VerifyTokenRepositorier interface {
	VerifyIdToken(context context.Context, token string) (dto.User, error)
}
