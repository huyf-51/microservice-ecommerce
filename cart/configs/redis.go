package configs

import (
	"os"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func RedisCache() {
	Rdb = redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_ADDR")})
}
