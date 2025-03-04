package usecase

import (
	"fmt"
	"log"
	"strconv"
	"time"
	"trivium/internal/domain/entity"
	"trivium/internal/domain/repositorier"
)

type MonitorCryptoCurrencies struct {
	cryptoCurrencyRepo repositorier.CryptoCurrencyRepositorier
	cryptoHistory      repositorier.CryptoHistoryRepository
}

func NewMonitorCryptoCurrencies(
	cryptoCurrencyRepo repositorier.CryptoCurrencyRepositorier,
	cryptoHistory repositorier.CryptoHistoryRepository) *MonitorCryptoCurrencies {
	return &MonitorCryptoCurrencies{
		cryptoCurrencyRepo: cryptoCurrencyRepo,
		cryptoHistory:      cryptoHistory,
	}
}

func (c *MonitorCryptoCurrencies) WatchCrypto() {
	cryptoEvent := NewCryptoWatchEventUseCase()

	var cryptoCurrencies, err = c.cryptoCurrencyRepo.FindAll()
	if err != nil {
		log.Fatal(err)
	}
	cryptos := make([]string, len(cryptoCurrencies))
	for i, crypto := range cryptoCurrencies {
		cryptos[i] = crypto.Symbol
	}

	dataChannel := cryptoEvent.WatchEvent(cryptos)

	go func() {
		for data := range dataChannel {
			price, err := strconv.ParseFloat(data.Price, 64)
			if err != nil {
				fmt.Printf("Erro ao converter preço: %v", err)
			}
			c.cryptoHistory.Save(entity.CryptoHistory{
				Name:      data.Symbol,
				Price:     price,
				Symbol:    data.Symbol,
				CreatedAt: time.Now(),
			})
			// log.Printf("Criptomoeda: %s, Preço: %s, Volume: %s", data.Symbol, data.Price, data.Volume)
		}
	}()
}
