package cmd

import (
	"trivium/internal/domain/repositorier"
	"trivium/internal/infra/http/repository"
	presentation_repositorier "trivium/internal/presentation/repositorier"

	"github.com/google/wire"
)

func NewFirebaseRepository() repositorier.AuthRepositorier {
	return repository.NewAuthRepository()
}

func NewHttpRepository() presentation_repositorier.HttpRepositorier {
	return repository.NewHttpRepository()
}

func NewVerifyTokenRepository() presentation_repositorier.VerifyTokenRepositorier {
	repo, err := repository.NewVerifyTokenRepository()
	if err != nil {
		panic(err)
	}
	return repo
}

var InfraModule = wire.NewSet(
	NewFirebaseRepository,
	NewHttpRepository,
	NewVerifyTokenRepository,
)
