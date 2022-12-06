package config

import (
	"github.com/go-redis/redis/v9"
	"os"
	"strconv"
)

func InitRedis() *redis.Client {
	RedisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       RedisDB,
		Network:  "tcp",
	})
}
