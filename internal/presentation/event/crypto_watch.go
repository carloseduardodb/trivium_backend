package event

import (
	"trivium/internal/domain/repositorier"
	"trivium/internal/infra/websocket/binance/repository"
)

type CryptoWatchEvent struct {
	Repository repositorier.CryptoStatusRepository
}

func NewCryptoWatchEvent() *CryptoWatchEvent {
	return &CryptoWatchEvent{
		Repository: repository.NewCryptoRepository(),
	}
}

func (c *CryptoWatchEvent) WatchEvent(cryptos []string) <-chan repositorier.CryptoData {
	return c.Repository.StreamCryptoData(cryptos)
}
