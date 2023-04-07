package connector

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"-"`
}

func ConnectPostgres(cfg *PostgresConfig) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := sqlx.Connect("postgres", connStr)

	if err != nil {
		return nil, fmt.Errorf("postgres error: %w", err)
	}
	return db, nil
}
