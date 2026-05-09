package usecase_test

import (
	"fmt"
	"testing"
	"time"

	"trivium/internal/domain/entity"
	"trivium/internal/domain/usecase"
	"trivium/internal/presentation/dto"
)

// Mock profit take repository
type mockProfitTakeRepo struct {
	profitTakes []entity.ProfitTake
	nextID      int64
}

func newMockProfitTakeRepo() *mockProfitTakeRepo {
	return &mockProfitTakeRepo{
		profitTakes: []entity.ProfitTake{},
		nextID:      1,
	}
}

func (m *mockProfitTakeRepo) Save(pt entity.ProfitTake) (entity.ProfitTake, error) {
	pt.ID = m.nextID
	m.nextID++
	m.profitTakes = append(m.profitTakes, pt)
	return pt, nil
}

func (m *mockProfitTakeRepo) FindById(id int64) (entity.ProfitTake, error) {
	for _, pt := range m.profitTakes {
		if pt.ID == id {
			return pt, nil
		}
	}
	return entity.ProfitTake{}, fmt.Errorf("not found")
}

func (m *mockProfitTakeRepo) FindAll() ([]entity.ProfitTake, error) {
	return m.profitTakes, nil
}

func (m *mockProfitTakeRepo) FindByPositionId(positionId int64) ([]entity.ProfitTake, error) {
	var result []entity.ProfitTake
	for _, pt := range m.profitTakes {
		if pt.Position == positionId {
			result = append(result, pt)
		}
	}
	return result, nil
}

func (m *mockProfitTakeRepo) Update(pt entity.ProfitTake) (entity.ProfitTake, error) {
	for i, p := range m.profitTakes {
		if p.ID == pt.ID {
			m.profitTakes[i] = pt
			return pt, nil
		}
	}
	return entity.ProfitTake{}, fmt.Errorf("not found")
}

func (m *mockProfitTakeRepo) Delete(id int64) error {
	for i, pt := range m.profitTakes {
		if pt.ID == id {
			m.profitTakes = append(m.profitTakes[:i], m.profitTakes[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("not found")
}

func TestProfitTakeUseCase_Create_Success(t *testing.T) {
	posRepo := newMockPositionRepo()
	ptRepo := newMockProfitTakeRepo()

	// Create a position first
	posRepo.Save(entity.Position{
		UserID:         1,
		CryptoCurrency: 1,
		Quantity:       0.5,
		PurchasePrice:  50000.0,
		InvestedAmount: 25000.0,
		Status:         "active",
	})

	uc := usecase.NewProfitTakeUseCase(ptRepo, posRepo)

	input := &dto.CreateProfitTake{
		PositionID:      1,
		AmountWithdrawn: 5000.0,
		PriceAtWithdraw: 60000.0,
		WithdrawDate:    time.Now(),
	}

	result, err := uc.Create(1, input)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.ID != 1 {
		t.Errorf("expected ID 1, got %d", result.ID)
	}
	if result.AmountWithdrawn != 5000.0 {
		t.Errorf("expected amount 5000, got %f", result.AmountWithdrawn)
	}
}

func TestProfitTakeUseCase_Create_ClosedPosition(t *testing.T) {
	posRepo := newMockPositionRepo()
	ptRepo := newMockProfitTakeRepo()

	posRepo.Save(entity.Position{
		UserID:         1,
		CryptoCurrency: 1,
		Quantity:       0.5,
		PurchasePrice:  50000.0,
		InvestedAmount: 25000.0,
		Status:         "closed",
	})

	uc := usecase.NewProfitTakeUseCase(ptRepo, posRepo)

	input := &dto.CreateProfitTake{
		PositionID:      1,
		AmountWithdrawn: 5000.0,
		PriceAtWithdraw: 60000.0,
		WithdrawDate:    time.Now(),
	}

	_, err := uc.Create(1, input)
	if err == nil {
		t.Fatal("expected error for closed position")
	}
}

func TestProfitTakeUseCase_Create_Unauthorized(t *testing.T) {
	posRepo := newMockPositionRepo()
	ptRepo := newMockProfitTakeRepo()

	posRepo.Save(entity.Position{
		UserID:         1,
		CryptoCurrency: 1,
		Quantity:       0.5,
		PurchasePrice:  50000.0,
		InvestedAmount: 25000.0,
		Status:         "active",
	})

	uc := usecase.NewProfitTakeUseCase(ptRepo, posRepo)

	input := &dto.CreateProfitTake{
		PositionID:      1,
		AmountWithdrawn: 5000.0,
		PriceAtWithdraw: 60000.0,
		WithdrawDate:    time.Now(),
	}

	_, err := uc.Create(999, input) // different user
	if err == nil {
		t.Fatal("expected unauthorized error")
	}
}

func TestProfitTakeUseCase_Create_ExceedsValue(t *testing.T) {
	posRepo := newMockPositionRepo()
	ptRepo := newMockProfitTakeRepo()

	posRepo.Save(entity.Position{
		UserID:         1,
		CryptoCurrency: 1,
		Quantity:       0.5,
		PurchasePrice:  50000.0,
		InvestedAmount: 25000.0,
		Status:         "active",
	})

	uc := usecase.NewProfitTakeUseCase(ptRepo, posRepo)

	input := &dto.CreateProfitTake{
		PositionID:      1,
		AmountWithdrawn: 999999.0, // exceeds position value
		PriceAtWithdraw: 60000.0,
		WithdrawDate:    time.Now(),
	}

	_, err := uc.Create(1, input)
	if err == nil {
		t.Fatal("expected error for exceeding position value")
	}
}
