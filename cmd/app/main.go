package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"trivium/internal"
	influxConnection "trivium/internal/infra/database/influx/connection"
	"trivium/internal/infra/database/postgres/connection"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	validateEnv()

	app, err := initializeApp()
	if err != nil {
		log.Fatalf("Error initializing application: %v", err)
	}

	internal.Bootstrap()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.Server.Start(); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	log.Println("Application started successfully")

	<-quit
	log.Println("Shutting down server...")

	connection.CloseDB()
	influxConnection.CloseDB()
	log.Println("Server stopped gracefully")
}

func validateEnv() {
	required := []string{
		"DATABASE_HOST",
		"DATABASE_PORT",
		"DATABASE_USER",
		"DATABASE_PASSWORD",
		"DATABASE_NAME",
	}

	for _, key := range required {
		if os.Getenv(key) == "" {
			log.Fatalf("Required environment variable %s is not set", key)
		}
	}
}
