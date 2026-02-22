package cache

import (
	"context"
	"log"

	"movie-watchlist-backend/internal/config"
	"github.com/redis/go-redis/v9"
)

var Client *redis.Client
var Ctx = context.Background()

func ConnectRedis(cfg *config.Config) {
	Client = redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPass,
		DB:       0,
	})

	_, err := Client.Ping(Ctx).Result()
	if err != nil {
		log.Printf("Failed to connect to Redis: %v. Caching will gracefully fail or application might panic depending on usage.", err)
	} else {
		log.Println("Redis connection established")
	}
}
