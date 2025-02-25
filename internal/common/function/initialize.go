package function

import (
	"trivium/internal/domain/usecase"
	"trivium/internal/infra/database/postgres/repository"
)

func Bootstrap() {
	repo := repository.NewCryptoCurrencyRepository()
	usecase.NewMonitorCryptoCurrencies(repo).WatchCrypto()
}
