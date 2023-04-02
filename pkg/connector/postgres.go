package connector

import (
	"database/sql"
	"fmt"
	"github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
	"strconv"
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

func GetPostgresConnector(cfg *PostgresConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		cfg.User,
		cfg.Database,
		cfg.Password,
		cfg.Host,
		strconv.FormatUint(uint64(cfg.Port), 10))

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// TODO: Нужно оптимизировать динамически максимальное число открытых соеднинений (и idle - тоже)
	db.SetMaxOpenConns(10)
	//db.SetMaxIdleConns()
	//db.SetConnMaxIdleTime()
	return db, nil
}

func GetSqlxConnector(db *sql.DB, driverName string) *sqlx.DB {
	return sqlx.NewDb(db, driverName)
}
