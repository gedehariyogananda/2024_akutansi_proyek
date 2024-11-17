package Config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

// global variable
var RedisClient *redis.Client
var Ctx = context.Background()

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})

	// checked connection ping
	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("log: failed connect to redis: %v", err)
	}

	log.Println("log: Redis connected!")
}

func SetToRedis(key string, value interface{}, expiration time.Duration) error {
	return RedisClient.Set(Ctx, key, value, expiration).Err()
}

func GetFromRedis(key string) (string, error) {
	return RedisClient.Get(Ctx, key).Result()
}

func DeleteFromRedis(key string) error {
	return RedisClient.Del(Ctx, key).Err()
}
