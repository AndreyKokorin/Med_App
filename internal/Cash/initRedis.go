package Cash

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log/slog"
)

var RedisClient *redis.Client

func InitRedis() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	var сtx = context.Background()

	if _, err := RedisClient.Ping(сtx).Result(); err != nil {
		return err
	}

	slog.Info("Подключено к Redis")
	return nil
}
