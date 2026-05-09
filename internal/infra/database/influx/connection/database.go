package connection

import (
	"log"
	"os"
	"sync"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var instance influxdb2.Client
var once sync.Once

func GetDB() influxdb2.Client {
	once.Do(func() {
		url := os.Getenv("INFLUXDB_URL")
		token := os.Getenv("INFLUXDB_TOKEN")

		if url == "" || token == "" {
			log.Println("Warning: INFLUXDB_URL or INFLUXDB_TOKEN not set")
		}

		instance = influxdb2.NewClient(url, token)
	})
	return instance
}

func CloseDB() {
	if instance != nil {
		instance.Close()
		instance = nil
	}
}
