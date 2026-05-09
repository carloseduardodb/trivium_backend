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

type VolumeRepositoryImpl struct{}

func NewVolumeRepository() repositorier.VolumeRepository {
	return &VolumeRepositoryImpl{}
}

func (r *VolumeRepositoryImpl) Save(volume entity.Volume) (entity.Volume, error) {
	client := connection.GetDB()
	writeAPI := client.WriteAPIBlocking(os.Getenv("INFLUXDB_ORG"), os.Getenv("INFLUXDB_BUCKET"))

	point := influxdb2.NewPoint("volume",
		map[string]string{"symbol": volume.Symbol},
		map[string]interface{}{
			"name":  volume.Name,
			"price": volume.Price,
		},
		volume.CreatedAt,
	)

	err := writeAPI.WritePoint(context.Background(), point)
	if err != nil {
		return entity.Volume{}, err
	}

	return volume, nil
}

func (r *VolumeRepositoryImpl) FindBySymbol(symbol string) ([]entity.Volume, error) {
	client := connection.GetDB()
	queryAPI := client.QueryAPI(os.Getenv("INFLUXDB_ORG"))

	query := fmt.Sprintf(`
		from(bucket: "%s")
		|> range(start: -30d)
		|> filter(fn: (r) => r._measurement == "volume" and r.symbol == "%s")
		|> filter(fn: (r) => r._field == "price")
		|> sort(columns: ["_time"], desc: true)
	`, os.Getenv("INFLUXDB_BUCKET"), symbol)

	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	var volumeList []entity.Volume
	for result.Next() {
		price, _ := result.Record().Value().(float64)
		volSymbol, _ := result.Record().ValueByKey("symbol").(string)
		if volSymbol == "" {
			volSymbol = symbol
		}

		volumeList = append(volumeList, entity.Volume{
			Name:      volSymbol,
			Price:     price,
			Symbol:    volSymbol,
			CreatedAt: result.Record().Time(),
		})
	}

	if result.Err() != nil {
		return nil, result.Err()
	}

	return volumeList, nil
}

func (r *VolumeRepositoryImpl) FindAll() ([]entity.Volume, error) {
	client := connection.GetDB()
	queryAPI := client.QueryAPI(os.Getenv("INFLUXDB_ORG"))

	query := fmt.Sprintf(`
		from(bucket: "%s")
		|> range(start: -30d)
		|> filter(fn: (r) => r._measurement == "volume")
		|> filter(fn: (r) => r._field == "price")
		|> sort(columns: ["_time"], desc: true)
	`, os.Getenv("INFLUXDB_BUCKET"))

	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	var volumeList []entity.Volume
	for result.Next() {
		price, _ := result.Record().Value().(float64)
		volSymbol, _ := result.Record().ValueByKey("symbol").(string)

		volumeList = append(volumeList, entity.Volume{
			Name:      volSymbol,
			Price:     price,
			Symbol:    volSymbol,
			CreatedAt: result.Record().Time(),
		})
	}

	if result.Err() != nil {
		return nil, result.Err()
	}

	return volumeList, nil
}
