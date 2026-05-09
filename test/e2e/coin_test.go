package e2e

import (
	"net/http"
	"os"
	"testing"
)

func TestGetCryptoCurrencies_Unauthorized(t *testing.T) {
	resp, err := http.Get(os.Getenv("HOST") + ":" + os.Getenv("PORT") + "/cryptocurrencies")
	if err != nil {
		t.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status 401 Unauthorized, got %d", resp.StatusCode)
	}
}
