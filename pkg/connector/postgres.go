package connector

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sethvargo/go-retry"
	"runtime"
	"strconv"
	"time"
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

func ConnectPostgresPgx(cfg *PostgresConfig) (*pgx.Conn, error) {
	connCfg := pgx.ConnConfig{
		Host:     cfg.Host,
		Port:     cfg.Port,
		Database: cfg.Database,
		User:     cfg.User,
		Password: cfg.Password,
	}
	db, err := pgx.Connect(connCfg)

	if err != nil {
		return nil, fmt.Errorf("postgres error: %w", err)
	}
	return db, nil

}

func GetPostgresConnector(cfg *PostgresConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		cfg.User,
		cfg.Database,
		cfg.Password,
		cfg.Host,
		strconv.FormatUint(uint64(cfg.Port), 10))

	fmt.Println(dsn)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	if err := retry.Fibonacci(ctx, 1*time.Second, func(ctx context.Context) error {
		if err := db.PingContext(ctx); err != nil {
			// This marks the error as retryable
			return retry.RetryableError(err)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	// Управление открытыми соединениями
	// Использование пула коннектов
	const OPEN_CONS_FACTOR = 3
	var cons = runtime.NumCPU()
	db.SetMaxOpenConns(cons * OPEN_CONS_FACTOR)
	db.SetMaxIdleConns(cons)
	db.SetConnMaxIdleTime(10 * time.Second)
	db.SetConnMaxLifetime(0)

	return db, nil
}

func GetSqlxConnector(db *sql.DB, driverName string) *sqlx.DB {
	return sqlx.NewDb(db, driverName)
}
