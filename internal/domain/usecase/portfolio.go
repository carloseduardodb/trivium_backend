package usecase

import (
	"trivium/internal/domain/entity"
	"trivium/internal/domain/repositorier"
)

type PortfolioSummary struct {
	TotalInvested    float64           `json:"total_invested"`
	ActivePositions  int               `json:"active_positions"`
	ClosedPositions  int               `json:"closed_positions"`
	Positions        []PositionSummary `json:"positions"`
}

type PositionSummary struct {
	Position     entity.Position  `json:"position"`
	CurrentPrice float64          `json:"current_price"`
	CurrentValue float64          `json:"current_value"`
	ProfitLoss   float64          `json:"profit_loss"`
	ProfitPct    float64          `json:"profit_pct"`
	ProfitTakes  []entity.ProfitTake `json:"profit_takes"`
}

type PortfolioUseCase struct {
	positionRepo    repositorier.PositionRepositorier
	profitTakeRepo  repositorier.ProfitTakeRepositorier
	cryptoHistoryRepo repositorier.CryptoHistoryRepository
	cryptoRepo      repositorier.CryptoCurrencyRepositorier
}

func NewPortfolioUseCase(
	positionRepo repositorier.PositionRepositorier,
	profitTakeRepo repositorier.ProfitTakeRepositorier,
	cryptoHistoryRepo repositorier.CryptoHistoryRepository,
	cryptoRepo repositorier.CryptoCurrencyRepositorier,
) *PortfolioUseCase {
	return &PortfolioUseCase{
		positionRepo:    positionRepo,
		profitTakeRepo:  profitTakeRepo,
		cryptoHistoryRepo: cryptoHistoryRepo,
		cryptoRepo:      cryptoRepo,
	}
}

func (p *PortfolioUseCase) GetPortfolio(userId int64) (*PortfolioSummary, error) {
	positions, err := p.positionRepo.FindByUserId(userId)
	if err != nil {
		return nil, err
	}

	latestPrices, err := p.cryptoHistoryRepo.FindAll()
	if err != nil {
		return nil, err
	}

	priceMap := make(map[string]float64)
	for _, h := range latestPrices {
		if _, exists := priceMap[h.Symbol]; !exists {
			priceMap[h.Symbol] = h.Price
		}
	}

	cryptos, err := p.cryptoRepo.FindAll()
	if err != nil {
		return nil, err
	}
	cryptoMap := make(map[int64]string)
	for _, c := range cryptos {
		cryptoMap[c.ID] = c.Symbol
	}

	summary := &PortfolioSummary{}
	for _, pos := range positions {
		if pos.Status == "active" {
			summary.ActivePositions++
		} else {
			summary.ClosedPositions++
		}
		summary.TotalInvested += pos.InvestedAmount

		profitTakes, _ := p.profitTakeRepo.FindByPositionId(pos.ID)

		symbol := cryptoMap[pos.CryptoCurrency]
		currentPrice := priceMap[symbol]
		currentValue := pos.Quantity * currentPrice
		profitLoss := currentValue - pos.InvestedAmount
		profitPct := 0.0
		if pos.InvestedAmount > 0 {
			profitPct = (profitLoss / pos.InvestedAmount) * 100
		}

		summary.Positions = append(summary.Positions, PositionSummary{
			Position:     pos,
			CurrentPrice: currentPrice,
			CurrentValue: currentValue,
			ProfitLoss:   profitLoss,
			ProfitPct:    profitPct,
			ProfitTakes:  profitTakes,
		})
	}

	return summary, nil
}
