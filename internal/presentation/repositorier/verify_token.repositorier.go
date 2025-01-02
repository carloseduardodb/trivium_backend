package presentation_repositorier

import (
	"context"
	"trivium/internal/common/types"
)

type VerifyTokenRepositorier interface {
	VerifyIdToken(context context.Context, token string) (types.User, error)
}
