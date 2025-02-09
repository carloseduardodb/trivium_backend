package e2e

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	m.Run()
}
