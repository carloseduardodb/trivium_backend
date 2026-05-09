package usecase

import (
	"fmt"
	"log"

	"trivium/internal/domain/entity"
	"trivium/internal/domain/repositorier"
)

type PriceAlertUseCase struct {
	alertRepo  repositorier.PriceAlertRepositorier
	cryptoRepo repositorier.CryptoCurrencyRepositorier
}

func NewPriceAlertUseCase(
	alertRepo repositorier.PriceAlertRepositorier,
	cryptoRepo repositorier.CryptoCurrencyRepositorier,
) *PriceAlertUseCase {
	return &PriceAlertUseCase{
		alertRepo:  alertRepo,
		cryptoRepo: cryptoRepo,
	}
}

type CreateAlertInput struct {
	CryptoCurrency int64   `json:"crypto_currency"`
	TargetPrice    float64 `json:"target_price"`
	Direction      string  `json:"direction"` // "above" or "below"
}

func (p *PriceAlertUseCase) Create(userId int64, input *CreateAlertInput) (*entity.PriceAlert, error) {
	if input.TargetPrice <= 0 {
		return nil, fmt.Errorf("target_price must be greater than zero")
	}
	if input.Direction != "above" && input.Direction != "below" {
		return nil, fmt.Errorf("direction must be 'above' or 'below'")
	}

	crypto, err := p.cryptoRepo.FindById(input.CryptoCurrency)
	if err != nil {
		return nil, fmt.Errorf("cryptocurrency not found: %w", err)
	}

	alert := entity.PriceAlert{
		UserID:         userId,
		CryptoCurrency: input.CryptoCurrency,
		Symbol:         crypto.Symbol,
		TargetPrice:    input.TargetPrice,
		Direction:      input.Direction,
	}

	saved, err := p.alertRepo.Save(alert)
	if err != nil {
		return nil, fmt.Errorf("error saving price alert: %w", err)
	}

	return &saved, nil
}

func (p *PriceAlertUseCase) FindByUserId(userId int64) ([]entity.PriceAlert, error) {
	return p.alertRepo.FindByUserId(userId)
}

func (p *PriceAlertUseCase) Delete(id int64, userId int64) error {
	alerts, err := p.alertRepo.FindByUserId(userId)
	if err != nil {
		return err
	}

	for _, alert := range alerts {
		if alert.ID == id {
			return p.alertRepo.Delete(id)
		}
	}

	return fmt.Errorf("alert not found or unauthorized")
}

func (p *PriceAlertUseCase) CheckAlerts(priceMap map[string]float64) []entity.PriceAlert {
	alerts, err := p.alertRepo.FindActive()
	if err != nil {
		log.Printf("Error fetching active alerts: %v", err)
		return nil
	}

	var triggered []entity.PriceAlert
	for _, alert := range alerts {
		currentPrice, exists := priceMap[alert.Symbol]
		if !exists {
			continue
		}

		shouldTrigger := false
		if alert.Direction == "above" && currentPrice >= alert.TargetPrice {
			shouldTrigger = true
		} else if alert.Direction == "below" && currentPrice <= alert.TargetPrice {
			shouldTrigger = true
		}

		if shouldTrigger {
			p.alertRepo.Deactivate(alert.ID)
			triggered = append(triggered, alert)
		}
	}

	return triggered
}
