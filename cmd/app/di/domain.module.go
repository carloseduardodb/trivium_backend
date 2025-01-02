package cmd

import (
	"trivium/internal/domain/usecase"

	"github.com/google/wire"
)

var DomainModule = wire.NewSet(
	usecase.NewAuthUseCase,
)
