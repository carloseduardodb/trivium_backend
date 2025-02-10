package connection

import (
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var instance influxdb2.Client

func GetDB() influxdb2.Client {
	if instance == nil {
		instance = influxdb2.NewClient(
			os.Getenv("INFLUXDB_URL"),
			os.Getenv("INFLUXDB_TOKEN"))
	}
	return instance
}
