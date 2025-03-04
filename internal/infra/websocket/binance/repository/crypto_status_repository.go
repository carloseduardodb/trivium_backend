package repository

import (
	"context"
	"log"
	"strconv"
	"sync"
	"time"

	"trivium/internal/domain/repositorier"

	"github.com/adshao/go-binance/v2"
)

type CryptoStatusRepositoryImpl struct{}

func NewCryptoRepository() repositorier.CryptoStatusRepository {
	return &CryptoStatusRepositoryImpl{}
}

func (r *CryptoStatusRepositoryImpl) Get24hVolumes(cryptos []string) (map[string]string, error) {
	client := binance.NewClient("", "")
	stats, err := client.NewListPriceChangeStatsService().Do(context.Background())
	if err != nil {
		return nil, err
	}

	volumes := make(map[string]string)
	for _, stat := range stats {
		for _, crypto := range cryptos {
			if stat.Symbol == crypto {
				volumes[stat.Symbol] = stat.Volume
			}
		}
	}
	return volumes, nil
}

func (r *CryptoStatusRepositoryImpl) StreamCryptoData(cryptos []string) <-chan repositorier.CryptoData {
	dataStream := make(chan repositorier.CryptoData)
	ticker := time.NewTicker(1 * time.Second * 10)
	var mu sync.Mutex

	cryptoMap := make(map[string]repositorier.CryptoData)

	wsHandler := func(event *binance.WsAggTradeEvent) {
		mu.Lock()
		defer mu.Unlock()

		price, err := strconv.ParseFloat(event.Price, 64)
		if err != nil {
			log.Println("Erro ao converter preço:", err)
			return
		}

		data := repositorier.CryptoData{
			Symbol: event.Symbol,
			Price:  strconv.FormatFloat(price, 'f', 8, 64),
			Volume: event.Quantity,
		}

		cryptoMap[event.Symbol] = data
	}

	errHandler := func(err error) {
		log.Println("Erro no WebSocket:", err)
	}

	for _, crypto := range cryptos {
		symbol := crypto + "USDT"

		_, _, err := binance.WsAggTradeServe(symbol, wsHandler, errHandler)
		if err != nil {
			log.Printf("Erro ao conectar ao WebSocket da Binance para %s: %v\n", symbol, err)
		}
	}

	go func() {
		defer ticker.Stop()
		defer close(dataStream)
		for range ticker.C {
			mu.Lock()
			for _, data := range cryptoMap {
				dataStream <- data
			}
			mu.Unlock()
		}
	}()

	log.Println("Conectado ao WebSocket da Binance para", cryptos)

	return dataStream
}
