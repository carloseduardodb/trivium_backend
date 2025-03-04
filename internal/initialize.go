package internal

import (
	"trivium/internal/domain/usecase"
	influxRepo "trivium/internal/infra/database/influx/repository"
	pgRepo "trivium/internal/infra/database/postgres/repository"
)

func Bootstrap() {
	cryptoCurrencyRepo := pgRepo.NewCryptoCurrencyRepository()
	cryptoHistoryRepo := influxRepo.NewCryptoHistoryRepository()
	usecase.NewMonitorCryptoCurrencies(cryptoCurrencyRepo, cryptoHistoryRepo).WatchCrypto()
}
