package connection

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var instance *sqlx.DB

var once sync.Once

func GetDB() *sqlx.DB {
	var (
		host     = os.Getenv("DATABASE_HOST")
		port     = os.Getenv("DATABASE_PORT")
		user     = os.Getenv("DATABASE_USER")
		password = os.Getenv("DATABASE_PASSWORD")
		dbname   = os.Getenv("DATABASE_NAME")
	)

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Fatal("Error when trying to get database connection: missing environment variables")
	}

	once.Do(func() {
		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
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

func ExecuteQuery(query string, args ...any) (*sqlx.Rows, error) {
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
