package main

import (
	"log"
	"trivium/internal"

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

	internal.Bootstrap()

	if err := app.Server.Start(); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
