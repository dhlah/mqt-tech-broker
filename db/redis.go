package db

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client
var Ctx = context.Background()

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "141.11.25.130:6379",
		Password: "rwSz4fcHaSShvrneU5aDQHZyX",
	})
}
