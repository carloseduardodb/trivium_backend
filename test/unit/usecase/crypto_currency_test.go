package usecase_test

import (
	"fmt"
	"testing"

	"trivium/internal/domain/entity"
	"trivium/internal/domain/usecase"
	"trivium/internal/presentation/dto"
)

// Mock repository
type mockCryptoCurrencyRepo struct {
	cryptos []entity.CryptoCurrency
	nextID  int64
}

func newMockCryptoRepo() *mockCryptoCurrencyRepo {
	return &mockCryptoCurrencyRepo{
		cryptos: []entity.CryptoCurrency{},
		nextID:  1,
	}
}

func (m *mockCryptoCurrencyRepo) Save(crypto entity.CryptoCurrency) (entity.CryptoCurrency, error) {
	crypto.ID = m.nextID
	m.nextID++
	m.cryptos = append(m.cryptos, crypto)
	return crypto, nil
}

func (m *mockCryptoCurrencyRepo) FindById(id int64) (entity.CryptoCurrency, error) {
	for _, c := range m.cryptos {
		if c.ID == id {
			return c, nil
		}
	}
	return entity.CryptoCurrency{}, fmt.Errorf("not found")
}

func (m *mockCryptoCurrencyRepo) FindAll() ([]entity.CryptoCurrency, error) {
	return m.cryptos, nil
}

func (m *mockCryptoCurrencyRepo) Update(crypto entity.CryptoCurrency) (entity.CryptoCurrency, error) {
	for i, c := range m.cryptos {
		if c.ID == crypto.ID {
			m.cryptos[i] = crypto
			return crypto, nil
		}
	}
	return entity.CryptoCurrency{}, fmt.Errorf("not found")
}

func (m *mockCryptoCurrencyRepo) Delete(id int64) error {
	for i, c := range m.cryptos {
		if c.ID == id {
			m.cryptos = append(m.cryptos[:i], m.cryptos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("not found")
}

func TestCryptoCurrencyUseCase_Create_Success(t *testing.T) {
	repo := newMockCryptoRepo()
	uc := usecase.NewCryptoCurrencyUseCase(repo)

	input := &dto.CreateCryptoCurrency{
		Name:   "Bitcoin",
		Symbol: "BTC",
	}

	result, err := uc.Create(input)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.ID != 1 {
		t.Errorf("expected ID 1, got %d", result.ID)
	}
	if result.Name != "Bitcoin" {
		t.Errorf("expected name 'Bitcoin', got '%s'", result.Name)
	}
	if result.Symbol != "BTC" {
		t.Errorf("expected symbol 'BTC', got '%s'", result.Symbol)
	}
}

func TestCryptoCurrencyUseCase_Create_EmptyName(t *testing.T) {
	repo := newMockCryptoRepo()
	uc := usecase.NewCryptoCurrencyUseCase(repo)

	input := &dto.CreateCryptoCurrency{
		Name:   "",
		Symbol: "BTC",
	}

	_, err := uc.Create(input)
	if err == nil {
		t.Fatal("expected error for empty name")
	}
}

func TestCryptoCurrencyUseCase_Create_EmptySymbol(t *testing.T) {
	repo := newMockCryptoRepo()
	uc := usecase.NewCryptoCurrencyUseCase(repo)

	input := &dto.CreateCryptoCurrency{
		Name:   "Bitcoin",
		Symbol: "",
	}

	_, err := uc.Create(input)
	if err == nil {
		t.Fatal("expected error for empty symbol")
	}
}

func TestCryptoCurrencyUseCase_FindAll(t *testing.T) {
	repo := newMockCryptoRepo()
	uc := usecase.NewCryptoCurrencyUseCase(repo)

	uc.Create(&dto.CreateCryptoCurrency{Name: "Bitcoin", Symbol: "BTC"})
	uc.Create(&dto.CreateCryptoCurrency{Name: "Ethereum", Symbol: "ETH"})

	result, err := uc.FindAll()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(result) != 2 {
		t.Errorf("expected 2 cryptos, got %d", len(result))
	}
}

func TestCryptoCurrencyUseCase_Delete(t *testing.T) {
	repo := newMockCryptoRepo()
	uc := usecase.NewCryptoCurrencyUseCase(repo)

	uc.Create(&dto.CreateCryptoCurrency{Name: "Bitcoin", Symbol: "BTC"})

	err := uc.Delete(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	result, _ := uc.FindAll()
	if len(result) != 0 {
		t.Errorf("expected 0 cryptos after delete, got %d", len(result))
	}
}
