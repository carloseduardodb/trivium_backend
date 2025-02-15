package repository

import (
	"context"
	"fmt"
	"os"

	"trivium/internal/domain/entity"
	"trivium/internal/domain/repositorier"
	"trivium/internal/infra/database/influx/connection"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type ProfitHistoryRepositoryImpl struct{}

func NewProfitHistoryRepository() repositorier.ProfitHistoryRepository {
	return &ProfitHistoryRepositoryImpl{}
}

func (r *ProfitHistoryRepositoryImpl) Save(profit entity.ProfitHistory) (entity.ProfitHistory, error) {
	client := connection.GetDB()
	writeAPI := client.WriteAPIBlocking(os.Getenv("INFLUXDB_ORG"), os.Getenv("INFLUXDB_BUCKET"))

	point := influxdb2.NewPoint("profit_history",
		map[string]string{
			"symbol":  profit.Symbol,
			"user_id": fmt.Sprintf("%d", profit.UserId),
		},
		map[string]interface{}{
			"crypto_price": profit.CryptoPrice,
			"profit":       profit.Profit,
		},
		profit.CreatedAt,
	)

	err := writeAPI.WritePoint(context.Background(), point)
	if err != nil {
		return entity.ProfitHistory{}, err
	}

	fmt.Println("ProfitHistory inserido com sucesso!")
	return profit, nil
}

func (r *ProfitHistoryRepositoryImpl) FindBySymbol(symbol string) ([]entity.ProfitHistory, error) {
	client := connection.GetDB()
	queryAPI := client.QueryAPI(os.Getenv("INFLUXDB_ORG"))

	query := fmt.Sprintf(`
		from(bucket: "%s") 
		|> range(start: -30d) 
		|> filter(fn: (r) => r._measurement == "profit_history" and r.symbol == "%s")
		|> sort(columns: ["_time"], desc: true)`, os.Getenv("INFLUXDB_BUCKET"), symbol)

	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	var profitList []entity.ProfitHistory
	for result.Next() {
		cryptoPrice, _ := result.Record().ValueByKey("crypto_price").(float64)
		profit, _ := result.Record().ValueByKey("profit").(float64)
		userIDStr, _ := result.Record().ValueByKey("user_id").(string)
		var userID int64
		fmt.Sscanf(userIDStr, "%d", &userID)

		profitList = append(profitList, entity.ProfitHistory{
			Symbol:      symbol,
			CryptoPrice: cryptoPrice,
			UserId:      userID,
			Profit:      profit,
			CreatedAt:   result.Record().Time(),
		})
	}

	if result.Err() != nil {
		return nil, result.Err()
	}

	return profitList, nil
}

func (r *ProfitHistoryRepositoryImpl) FindAll() ([]entity.ProfitHistory, error) {
	client := connection.GetDB()
	queryAPI := client.QueryAPI(os.Getenv("INFLUXDB_ORG"))

	query := fmt.Sprintf(`
		from(bucket: "%s") 
		|> range(start: -30d) 
		|> filter(fn: (r) => r._measurement == "profit_history")
		|> sort(columns: ["_time"], desc: true)`, os.Getenv("INFLUXDB_BUCKET"))

	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	var profitList []entity.ProfitHistory
	for result.Next() {
		cryptoPrice, _ := result.Record().ValueByKey("crypto_price").(float64)
		profit, _ := result.Record().ValueByKey("profit").(float64)
		symbol, _ := result.Record().ValueByKey("symbol").(string)
		userIDStr, _ := result.Record().ValueByKey("user_id").(string)
		var userID int64
		fmt.Sscanf(userIDStr, "%d", &userID)

		profitList = append(profitList, entity.ProfitHistory{
			Symbol:      symbol,
			CryptoPrice: cryptoPrice,
			UserId:      userID,
			Profit:      profit,
			CreatedAt:   result.Record().Time(),
		})
	}

	if result.Err() != nil {
		return nil, result.Err()
	}

	return profitList, nil
}
