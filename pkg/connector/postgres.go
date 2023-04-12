package connector

import (
	"database/sql"
	"fmt"
	"github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"strconv"
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
