package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"trivium/internal/infra/database/postgres/repository"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	command := flag.String("command", "up", "migration command (up/down/status/create)")
	name := flag.String("name", "", "name for new migration")
	migType := flag.String("type", "sql", "migration type (sql/go)")
	flag.Parse()

	var (
		host     = os.Getenv("DATABASE_HOST")
		port     = os.Getenv("DATABASE_PORT")
		user     = os.Getenv("DATABASE_USER")
		password = os.Getenv("DATABASE_PASSWORD")
		dbname   = os.Getenv("DATABASE_NAME")
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	migration, err := repository.NewMigration(dsn, "internal/infra/database/postgres/migration")
	if err != nil {
		log.Fatalf("Failed to initialize migration: %v", err)
	}
	defer migration.Close()

	seed, err := repository.NewMigration(dsn, "internal/infra/database/postgres/seed")
	if err != nil {
		log.Fatalf("Failed to initialize migration: %v", err)
	}
	defer seed.Close()

	switch *command {
	case "up":
		err = migration.Up()
	case "down":
		err = migration.Down()
	case "status":
		err = migration.Status()
	case "create:migration":
		if *name == "" {
			log.Fatal("Name is required for create command")
		}
		err = migration.Create(*name, *migType)
	case "create:seed":
		if *name == "" {
			log.Fatal("Name is required for create command")
		}
		err = seed.Create(*name, *migType)
	case "seed":
		err = seed.Up()
	default:
		log.Fatalf("Unknown command: %s", *command)
	}

	if err != nil {
		log.Fatalf("Migration %s failed: %v", *command, err)
	}
}
