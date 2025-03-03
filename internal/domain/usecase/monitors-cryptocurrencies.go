package usecase

import (
	"log"
	"trivium/internal/domain/repositorier"
)

type MonitorCryptoCurrencies struct {
	cryptoCurrencyRepo repositorier.CryptoCurrencyRepositorier
}

func NewMonitorCryptoCurrencies(cryptoCurrencyRepo repositorier.CryptoCurrencyRepositorier) *MonitorCryptoCurrencies {
	return &MonitorCryptoCurrencies{
		cryptoCurrencyRepo: cryptoCurrencyRepo,
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
			log.Printf("Criptomoeda: %s, Preço: %s, Volume: %s", data.Symbol, data.Price, data.Volume)
		}
	}()
}
