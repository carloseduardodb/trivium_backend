package cmd

import (
	"trivium/internal/domain/repositorier"
	influxRepo "trivium/internal/infra/database/influx/repository"
	httpRepo "trivium/internal/infra/http/repository"
	pgRepo "trivium/internal/infra/database/postgres/repository"
	presentation_repositorier "trivium/internal/presentation/repositorier"

	"github.com/google/wire"
)

func NewFirebaseRepository() repositorier.AuthRepositorier {
	return httpRepo.NewAuthRepository()
}

func NewHttpRepository() presentation_repositorier.HttpRepositorier {
	return httpRepo.NewHttpRepository()
}

func NewVerifyTokenRepository() presentation_repositorier.VerifyTokenRepositorier {
	repo, err := httpRepo.NewVerifyTokenRepository()
	if err != nil {
		panic(err)
	}
	return repo
}

func NewUserRepository() repositorier.UserRepositorier {
	return pgRepo.NewUserRepository()
}

func NewCryptoCurrencyRepository() repositorier.CryptoCurrencyRepositorier {
	return pgRepo.NewCryptoCurrencyRepository()
}

func NewPositionRepository() repositorier.PositionRepositorier {
	return pgRepo.NewPositionRepository()
}

func NewProfitTakeRepository() repositorier.ProfitTakeRepositorier {
	return pgRepo.NewProfitTakeRepository()
}

func NewPriceAlertRepository() repositorier.PriceAlertRepositorier {
	return pgRepo.NewPriceAlertRepository()
}

func NewCryptoHistoryRepository() repositorier.CryptoHistoryRepository {
	return influxRepo.NewCryptoHistoryRepository()
}

func NewVolumeRepository() repositorier.VolumeRepository {
	return influxRepo.NewVolumeRepository()
}

var InfraModule = wire.NewSet(
	NewFirebaseRepository,
	NewHttpRepository,
	NewVerifyTokenRepository,
	NewUserRepository,
	NewCryptoCurrencyRepository,
	NewPositionRepository,
	NewProfitTakeRepository,
	NewPriceAlertRepository,
	NewCryptoHistoryRepository,
	NewVolumeRepository,
)
