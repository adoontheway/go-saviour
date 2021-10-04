package service

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:7001",
		Password: "12345",
		DB:       0,
	})
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
}
