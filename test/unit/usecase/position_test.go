package usecase_test

import (
	"fmt"
	"testing"
	"time"

	"trivium/internal/domain/entity"
	"trivium/internal/domain/usecase"
	"trivium/internal/presentation/dto"
)

// Mock position repository
type mockPositionRepo struct {
	positions []entity.Position
	nextID    int64
}

func newMockPositionRepo() *mockPositionRepo {
	return &mockPositionRepo{
		positions: []entity.Position{},
		nextID:    1,
	}
}

func (m *mockPositionRepo) Save(pos entity.Position) (entity.Position, error) {
	pos.ID = m.nextID
	m.nextID++
	m.positions = append(m.positions, pos)
	return pos, nil
}

func (m *mockPositionRepo) FindById(id int64) (entity.Position, error) {
	for _, p := range m.positions {
		if p.ID == id {
			return p, nil
		}
	}
	return entity.Position{}, fmt.Errorf("not found")
}

func (m *mockPositionRepo) FindAll() ([]entity.Position, error) {
	return m.positions, nil
}

func (m *mockPositionRepo) FindByUserId(userId int64) ([]entity.Position, error) {
	var result []entity.Position
	for _, p := range m.positions {
		if p.UserID == userId {
			result = append(result, p)
		}
	}
	return result, nil
}

func (m *mockPositionRepo) Update(pos entity.Position) (entity.Position, error) {
	for i, p := range m.positions {
		if p.ID == pos.ID {
			m.positions[i] = pos
			return pos, nil
		}
	}
	return entity.Position{}, fmt.Errorf("not found")
}

func (m *mockPositionRepo) Delete(id int64) error {
	for i, p := range m.positions {
		if p.ID == id {
			m.positions = append(m.positions[:i], m.positions[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("not found")
}

func TestPositionUseCase_Create_Success(t *testing.T) {
	posRepo := newMockPositionRepo()
	cryptoRepo := newMockCryptoRepo()

	// Add a crypto first
	cryptoRepo.Save(entity.CryptoCurrency{Name: "Bitcoin", Symbol: "BTC"})

	uc := usecase.NewPositionUseCase(posRepo, cryptoRepo)

	input := &dto.CreatePosition{
		CryptoCurrency: 1,
		Quantity:       0.5,
		PurchasePrice:  50000.0,
		PurchaseDate:   time.Now(),
	}

	result, err := uc.Create(1, input)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.ID != 1 {
		t.Errorf("expected ID 1, got %d", result.ID)
	}
	if result.InvestedAmount != 25000.0 {
		t.Errorf("expected invested amount 25000, got %f", result.InvestedAmount)
	}
	if result.Status != "active" {
		t.Errorf("expected status 'active', got '%s'", result.Status)
	}
}

func TestPositionUseCase_Create_InvalidQuantity(t *testing.T) {
	posRepo := newMockPositionRepo()
	cryptoRepo := newMockCryptoRepo()
	cryptoRepo.Save(entity.CryptoCurrency{Name: "Bitcoin", Symbol: "BTC"})

	uc := usecase.NewPositionUseCase(posRepo, cryptoRepo)

	input := &dto.CreatePosition{
		CryptoCurrency: 1,
		Quantity:       0,
		PurchasePrice:  50000.0,
		PurchaseDate:   time.Now(),
	}

	_, err := uc.Create(1, input)
	if err == nil {
		t.Fatal("expected error for zero quantity")
	}
}

func TestPositionUseCase_Create_CryptoNotFound(t *testing.T) {
	posRepo := newMockPositionRepo()
	cryptoRepo := newMockCryptoRepo()

	uc := usecase.NewPositionUseCase(posRepo, cryptoRepo)

	input := &dto.CreatePosition{
		CryptoCurrency: 999,
		Quantity:       0.5,
		PurchasePrice:  50000.0,
		PurchaseDate:   time.Now(),
	}

	_, err := uc.Create(1, input)
	if err == nil {
		t.Fatal("expected error for non-existent crypto")
	}
}

func TestPositionUseCase_Close_Success(t *testing.T) {
	posRepo := newMockPositionRepo()
	cryptoRepo := newMockCryptoRepo()
	cryptoRepo.Save(entity.CryptoCurrency{Name: "Bitcoin", Symbol: "BTC"})

	uc := usecase.NewPositionUseCase(posRepo, cryptoRepo)

	uc.Create(1, &dto.CreatePosition{
		CryptoCurrency: 1,
		Quantity:       0.5,
		PurchasePrice:  50000.0,
		PurchaseDate:   time.Now(),
	})

	result, err := uc.Close(1, 1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.Status != "closed" {
		t.Errorf("expected status 'closed', got '%s'", result.Status)
	}
}

func TestPositionUseCase_Close_Unauthorized(t *testing.T) {
	posRepo := newMockPositionRepo()
	cryptoRepo := newMockCryptoRepo()
	cryptoRepo.Save(entity.CryptoCurrency{Name: "Bitcoin", Symbol: "BTC"})

	uc := usecase.NewPositionUseCase(posRepo, cryptoRepo)

	uc.Create(1, &dto.CreatePosition{
		CryptoCurrency: 1,
		Quantity:       0.5,
		PurchasePrice:  50000.0,
		PurchaseDate:   time.Now(),
	})

	_, err := uc.Close(1, 999) // different user
	if err == nil {
		t.Fatal("expected unauthorized error")
	}
}

func TestPositionUseCase_Close_AlreadyClosed(t *testing.T) {
	posRepo := newMockPositionRepo()
	cryptoRepo := newMockCryptoRepo()
	cryptoRepo.Save(entity.CryptoCurrency{Name: "Bitcoin", Symbol: "BTC"})

	uc := usecase.NewPositionUseCase(posRepo, cryptoRepo)

	uc.Create(1, &dto.CreatePosition{
		CryptoCurrency: 1,
		Quantity:       0.5,
		PurchasePrice:  50000.0,
		PurchaseDate:   time.Now(),
	})

	uc.Close(1, 1)
	_, err := uc.Close(1, 1)
	if err == nil {
		t.Fatal("expected error for already closed position")
	}
}
