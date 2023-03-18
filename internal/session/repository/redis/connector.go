package redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

func Connect(opts *redis.Options) (*redis.Client, error) {
	client := redis.NewClient(opts)

	err := client.Ping().Err()
	if err != nil {
		return nil, fmt.Errorf("redis error: %w", err)
	}
	return client, nil
}
