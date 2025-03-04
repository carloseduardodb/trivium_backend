package usecase

import (
	"trivium/internal/domain/repositorier"
	"trivium/internal/infra/websocket/binance/repository"
)

type CryptoVolumeUseCase struct {
	Repository repositorier.CryptoStatusRepository
}

func NewCryptoVolumeUseCase() *CryptoVolumeUseCase {
	return &CryptoVolumeUseCase{
		Repository: repository.NewCryptoRepository(),
	}
}

func (c *CryptoVolumeUseCase) Get24hVolumes(cryptos []string) (map[string]string, error) {
	return c.Repository.Get24hVolumes(cryptos)
}
