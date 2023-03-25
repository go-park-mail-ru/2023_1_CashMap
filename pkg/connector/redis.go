package connector

import (
	"fmt"
	"github.com/go-redis/redis"
)

type RedisConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	DB   int    `yaml:"db"`
	Pass string `yaml:"-"`
}

func ConnectRedis(cfg *RedisConfig) (*redis.Client, error) {
	opts := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		DB:       cfg.DB,
		Password: cfg.Pass,
	}
	client := redis.NewClient(opts)
	err := client.Ping().Err()
	if err != nil {
		return nil, fmt.Errorf("redis error: %w", err)
	}
	return client, nil
}
