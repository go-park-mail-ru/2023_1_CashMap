package connector

import (
	"fmt"
	"github.com/jackc/pgx"
)

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"-"`
}

func ConnectPostgres(cfg *PostgresConfig) (*pgx.Conn, error) {
	connCfg := pgx.ConnConfig{
		Host:     cfg.Host,
		Port:     cfg.Port,
		Database: cfg.Database,
		User:     cfg.User,
		Password: cfg.Password,
	}
	conn, err := pgx.Connect(connCfg)
	if err != nil {
		return nil, fmt.Errorf("postgres error: %w", err)
	}
	return conn, nil
}
