package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v9"
	"time"
)

type RedisCache struct {
	redis *redis.Client
}

func NewRedisCache(redis *redis.Client) *RedisCache {
	return &RedisCache{redis: redis}
}

func (r *RedisCache) Set(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	JsonData, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	ErrorRedis := r.redis.Set(ctx, key, JsonData, duration).Err()
	if ErrorRedis != nil {
		return ErrorRedis
	}
	return ErrorRedis
}

func (r *RedisCache) Get(ctx context.Context, key string) (interface{}, error) {
	valueRedis, ErrorRedis := r.redis.Get(ctx, key).Result()
	if ErrorRedis != nil {
		return nil, ErrorRedis
	}
	return valueRedis, nil
}
