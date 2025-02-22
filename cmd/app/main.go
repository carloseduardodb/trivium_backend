package main

import (
	"log"
	"trivium/internal/presentation/event"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app, err := initializeApp()
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}

	cryptoEvent := event.NewCryptoWatchEvent()

	cryptos := []string{"BTC"}

	dataChannel := cryptoEvent.WatchEvent(cryptos)

	go func() {
		for data := range dataChannel {
			log.Printf("Criptomoeda: %s, Preço: %s, Volume: %s", data.Symbol, data.Price, data.Volume)
		}
	}()

	if err := app.Server.Start(); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
