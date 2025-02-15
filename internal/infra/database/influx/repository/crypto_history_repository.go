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

type CryptoHistoryRepositoryImpl struct{}

func NewCryptoRepository() repositorier.CryptoHistoryRepository {
	return &CryptoHistoryRepositoryImpl{}
}

func (r *CryptoHistoryRepositoryImpl) Save(crypto entity.CryptoHistory) (entity.CryptoHistory, error) {
	client := connection.GetDB()
	writeAPI := client.WriteAPIBlocking(os.Getenv("INFLUXDB_ORG"), os.Getenv("INFLUXDB_BUCKET"))

	point := influxdb2.NewPoint("crypto_history",
		map[string]string{"symbol": crypto.Symbol},
		map[string]any{
			"name":  crypto.Name,
			"price": crypto.Price,
		},
		crypto.CreatedAt,
	)

	err := writeAPI.WritePoint(context.Background(), point)
	if err != nil {
		return entity.CryptoHistory{}, err
	}

	fmt.Println("CryptoHistory inserido com sucesso!")
	return crypto, nil
}

func (r *CryptoHistoryRepositoryImpl) FindBySymbol(symbol string) ([]entity.CryptoHistory, error) {
	client := connection.GetDB()
	queryAPI := client.QueryAPI(os.Getenv("INFLUXDB_ORG"))

	query := fmt.Sprintf(`
		from(bucket: "%s") 
		|> range(start: -30d) 
		|> filter(fn: (r) => r._measurement == "crypto_history" and r.symbol == "%s")
		|> sort(columns: ["_time"], desc: true)`, os.Getenv("INFLUXDB_BUCKET"), symbol)

	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	var cryptoList []entity.CryptoHistory
	for result.Next() {
		name, _ := result.Record().ValueByKey("name").(string)
		price, _ := result.Record().Value().(float64)

		cryptoList = append(cryptoList, entity.CryptoHistory{
			Name:      name,
			Price:     price,
			Symbol:    symbol,
			CreatedAt: result.Record().Time(),
		})
	}

	if result.Err() != nil {
		return nil, result.Err()
	}

	return cryptoList, nil
}

func (r *CryptoHistoryRepositoryImpl) FindAll() ([]entity.CryptoHistory, error) {
	client := connection.GetDB()
	queryAPI := client.QueryAPI(os.Getenv("INFLUXDB_ORG"))

	query := fmt.Sprintf(`
		from(bucket: "%s") 
		|> range(start: -30d) 
		|> filter(fn: (r) => r._measurement == "crypto_history")
		|> sort(columns: ["_time"], desc: true)`, os.Getenv("INFLUXDB_BUCKET"))

	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	var cryptoList []entity.CryptoHistory
	for result.Next() {
		name, _ := result.Record().ValueByKey("name").(string)
		price, _ := result.Record().Value().(float64)
		symbol, _ := result.Record().ValueByKey("symbol").(string)

		cryptoList = append(cryptoList, entity.CryptoHistory{
			Name:      name,
			Price:     price,
			Symbol:    symbol,
			CreatedAt: result.Record().Time(),
		})
	}

	if result.Err() != nil {
		return nil, result.Err()
	}

	return cryptoList, nil
}
