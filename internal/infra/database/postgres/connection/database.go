package connection

import (
	"fmt"
	"log"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var instance *sqlx.DB

var once sync.Once

const (
	host     = "localhost"
	port     = 5432
	user     = "yourusername"
	password = "yourpassword"
	dbname   = "yourdatabase"
)

func GetDB() *sqlx.DB {
	once.Do(func() {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)

		db, err := sqlx.Open("postgres", psqlInfo)
		if err != nil {
			log.Fatalf("Error when open connection with database: %v", err)
		}

		err = db.Ping()
		if err != nil {
			log.Fatalf("Error when verify connection with database: %v", err)
		}

		instance = db
	})

	return instance
}

func ExecuteQuery(query string, args ...interface{}) (*sqlx.Rows, error) {
	db := GetDB()
	defer CloseDB()

	rows, err := db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func CloseDB() {
	if instance != nil {
		err := instance.Close()
		if err != nil {
			log.Fatalf("Error when trying to close connection with database: %v", err)
		}
		instance = nil
	}
}
