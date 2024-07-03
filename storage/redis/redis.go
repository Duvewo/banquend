package redis

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func Open(connString string) (*redis.Client, error) {
	opt, err := redis.ParseURL(connString)

	if err != nil {
		return nil, fmt.Errorf("redis/open: to parse: %w", err)
	}

	return redis.NewClient(opt), nil
}

func OpenWithOptions(opt *redis.Options) *redis.Client {
	return redis.NewClient(opt)
}
