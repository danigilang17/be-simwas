package services

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Sesuaikan dengan konfigurasi Anda
	})
}

func SetRedisKey(key, value string, expiration int64) error {
	ctx := context.Background()
	duration := time.Duration(expiration) * time.Second
	return RedisClient.Set(ctx, key, value, duration).Err()
}

func GetRedisKey(key string) (string, error) {
	ctx := context.Background()
	return RedisClient.Get(ctx, key).Result()
}
