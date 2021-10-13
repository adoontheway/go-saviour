package service

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type red struct {
	client *redis.Client
}

func (r *red) Set(key string, value interface{}, expiration time.Duration) {
	r.client.Set(ctx, key, value, expiration)
}

func InitRedis() *red {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:7001",
		// Password: "12345",
		DB: 0,
	})
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	return &red{
		client: rdb,
	}
}
