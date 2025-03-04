package usecase

import (
	"log"
	"strconv"
	"time"
	"trivium/internal/domain/entity"
	"trivium/internal/domain/repositorier"
)

type MonitorCryptoCurrencies struct {
	cryptoCurrencyRepo repositorier.CryptoCurrencyRepositorier
	cryptoHistoryRepo  repositorier.CryptoHistoryRepository
	volumeRepo         repositorier.VolumeRepository
}

func NewMonitorCryptoCurrencies(
	cryptoCurrencyRepo repositorier.CryptoCurrencyRepositorier,
	cryptoHistory repositorier.CryptoHistoryRepository,
	volume repositorier.VolumeRepository,
) *MonitorCryptoCurrencies {
	return &MonitorCryptoCurrencies{
		cryptoCurrencyRepo: cryptoCurrencyRepo,
		cryptoHistoryRepo:  cryptoHistory,
		volumeRepo:         volume,
	}
}

func parseFloat(value string) float64 {
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func (c *MonitorCryptoCurrencies) WatchCrypto() {
	cryptoEvent := NewCryptoWatchEventUseCase()
	cryptoVolume := NewCryptoVolumeUseCase()

	var cryptoCurrencies, err = c.cryptoCurrencyRepo.FindAll()
	if err != nil {
		log.Fatal(err)
	}
	cryptos := make([]string, len(cryptoCurrencies))
	for i, crypto := range cryptoCurrencies {
		cryptos[i] = crypto.Symbol
	}

	volumes, err := cryptoVolume.Get24hVolumes(cryptos)
	if err != nil {
		log.Fatal(err)
	}

	dataChannel := cryptoEvent.WatchEvent(cryptos)

	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		for range ticker.C {
			volumes, err = cryptoVolume.Get24hVolumes(cryptos)
			if err != nil {
				log.Fatal(err)
			}

			for symbol, volume := range volumes {
				c.volumeRepo.Save(entity.Volume{
					Name:      symbol,
					Price:     parseFloat(volume),
					Symbol:    symbol,
					CreatedAt: time.Now(),
				})
			}
		}
	}()

	go func() {
		for data := range dataChannel {
			c.cryptoHistoryRepo.Save(entity.CryptoHistory{
				Name:      data.Symbol,
				Price:     parseFloat(data.Price),
				Symbol:    data.Symbol,
				CreatedAt: time.Now(),
			})
		}
	}()
}
