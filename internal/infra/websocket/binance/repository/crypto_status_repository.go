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
	dataStream := make(chan repositorier.CryptoData, 100)
	ticker := time.NewTicker(10 * time.Second)
	var mu sync.Mutex

	cryptoMap := make(map[string]repositorier.CryptoData)

	wsHandler := func(event *binance.WsAggTradeEvent) {
		mu.Lock()
		defer mu.Unlock()

		price, err := strconv.ParseFloat(event.Price, 64)
		if err != nil {
			log.Println("Error converting price:", err)
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
		log.Println("WebSocket error:", err)
	}

	for _, crypto := range cryptos {
		symbol := crypto + "USDT"
		go r.connectWithRetry(symbol, wsHandler, errHandler)
	}

	go func() {
		defer ticker.Stop()
		for range ticker.C {
			mu.Lock()
			for _, data := range cryptoMap {
				select {
				case dataStream <- data:
				default:
					// channel full, skip
				}
			}
			mu.Unlock()
		}
	}()

	log.Println("Connected to Binance WebSocket for", cryptos)

	return dataStream
}

func (r *CryptoStatusRepositoryImpl) connectWithRetry(symbol string, wsHandler func(*binance.WsAggTradeEvent), errHandler func(error)) {
	maxRetries := 10
	baseDelay := 1 * time.Second

	for attempt := 0; attempt < maxRetries; attempt++ {
		doneC, _, err := binance.WsAggTradeServe(symbol, wsHandler, errHandler)
		if err != nil {
			delay := baseDelay * time.Duration(1<<uint(attempt))
			if delay > 60*time.Second {
				delay = 60 * time.Second
			}
			log.Printf("WebSocket connection failed for %s (attempt %d/%d), retrying in %v: %v",
				symbol, attempt+1, maxRetries, delay, err)
			time.Sleep(delay)
			continue
		}

		log.Printf("WebSocket connected for %s", symbol)
		<-doneC
		log.Printf("WebSocket disconnected for %s, reconnecting...", symbol)
		attempt = 0
		time.Sleep(2 * time.Second)
	}

	log.Printf("WebSocket max retries reached for %s", symbol)
}
