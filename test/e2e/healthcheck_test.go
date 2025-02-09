package e2e

import (
	"net/http"
	"os"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	resp, err := http.Get(os.Getenv("HOST") + ":" + os.Getenv("PORT") + "/healthcheck")

	if err != nil {
		t.Fatalf("Erro ao fazer a requisição: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Esperado status 200 OK, mas obteve %d", resp.StatusCode)
	}
}
