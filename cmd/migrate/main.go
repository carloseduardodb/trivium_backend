package main

import (
	"flag"
	"log"
	"trivium/internal/infra/database/postgres/repository"
)

func main() {
	command := flag.String("command", "up", "migration command (up/down/status/create)")
	name := flag.String("name", "", "name for new migration")
	migType := flag.String("type", "sql", "migration type (sql/go)")
	flag.Parse()

	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=mydatabase sslmode=disable"
	migration, err := repository.NewMigration(dsn, "internal/infra/database/postgres/migration")
	if err != nil {
		log.Fatalf("Failed to initialize migration: %v", err)
	}
	defer migration.Close()

	switch *command {
	case "up":
		err = migration.Up()
	case "down":
		err = migration.Down()
	case "status":
		err = migration.Status()
	case "create":
		if *name == "" {
			log.Fatal("Name is required for create command")
		}
		err = migration.Create(*name, *migType)
	default:
		log.Fatalf("Unknown command: %s", *command)
	}

	if err != nil {
		log.Fatalf("Migration %s failed: %v", *command, err)
	}
}
