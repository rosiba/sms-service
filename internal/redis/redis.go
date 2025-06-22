package rds

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
)

type RedisRepository struct {
	rds *redis.Client
}

func NewRedisRepository() *RedisRepository {
	redisURL := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	c := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // Default DB used for simplicity
	})
	return &RedisRepository{
		rds: c,
	}
}

func (r *RedisRepository) SaveMessageTime(id string, sentAt time.Time) error {
	ctx := context.Background()
	if err := r.rds.Set(ctx, id, sentAt, 0).Err(); err != nil {
		return fmt.Errorf("failed to save message time: %w", err)
	}

	return nil
}
