package repository

import (
	"log"
	"time"

	"trivium/internal/domain/repositorier"

	"github.com/adshao/go-binance/v2"
)

type CryptoStatusRepositoryImpl struct{}

func NewCryptoRepository() repositorier.CryptoStatusRepository {
	return &CryptoStatusRepositoryImpl{}
}

func (r *CryptoStatusRepositoryImpl) StreamCryptoData(cryptos []string) <-chan repositorier.CryptoData {
	dataStream := make(chan repositorier.CryptoData)
	ticker := time.NewTicker(1 * time.Minute)

	cryptoMap := make(map[string]repositorier.CryptoData)

	wsHandler := func(event *binance.WsAggTradeEvent) {
		data := repositorier.CryptoData{
			Symbol: event.Symbol,
			Price:  event.Price,
			Volume: event.Quantity,
		}

		cryptoMap[event.Symbol] = data
	}

	errHandler := func(err error) {
		log.Println("Erro no WebSocket:", err)
	}

	for _, crypto := range cryptos {
		symbol := crypto + "usdt"
		_, _, err := binance.WsAggTradeServe(symbol, wsHandler, errHandler)
		if err != nil {
			log.Printf("Erro ao conectar ao WebSocket da Binance para %s: %v\n", symbol, err)
		}
	}

	go func() {
		defer ticker.Stop()
		defer close(dataStream)
		for range ticker.C {
			for _, data := range cryptoMap {
				dataStream <- data
			}
		}
	}()

	return dataStream
}
