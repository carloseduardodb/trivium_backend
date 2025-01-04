package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

type Migration struct {
	db   *sql.DB
	path string
}

func NewMigration(dsn string, migrationsPath string) (*Migration, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		return nil, fmt.Errorf("failed to set dialect: %v", err)
	}

	return &Migration{
		db:   db,
		path: migrationsPath,
	}, nil
}

func (m *Migration) Up() error {
	return goose.Up(m.db, m.path)
}

func (m *Migration) Down() error {
	return goose.Down(m.db, m.path)
}

func (m *Migration) Status() error {
	return goose.Status(m.db, m.path)
}

func (m *Migration) Create(name, migrationType string) error {
	return goose.Create(m.db, m.path, name, migrationType)
}

func (m *Migration) Close() error {
	return m.db.Close()
}
