package usecase

import (
	"trivium/internal/domain/repositorier"
	"trivium/internal/infra/websocket/binance/repository"
)

type CryptoWatchEventUseCase struct {
	Repository repositorier.CryptoStatusRepository
}

func NewCryptoWatchEventUseCase() *CryptoWatchEventUseCase {
	return &CryptoWatchEventUseCase{
		Repository: repository.NewCryptoRepository(),
	}
}

func (c *CryptoWatchEventUseCase) WatchEvent(cryptos []string) <-chan repositorier.CryptoData {
	return c.Repository.StreamCryptoData(cryptos)
}
