package redis

import (
	"backend/config"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
	ctx    context.Context
}

func NewCache(cfg *config.RedisConfig) *Cache {
	return &Cache{
		client: redis.NewClient(&redis.Options{
			Addr:     cfg.Addr,
			Password: cfg.Password,
			DB:       cfg.DB,
		}),
		ctx: context.Background(),
	}
}

func (c *Cache) Set(key string, value interface{}, expiration time.Duration) error {
	return c.client.Set(c.ctx, key, value, expiration).Err()
}

func (c *Cache) Get(key string) (string, error) {
	return c.client.Get(c.ctx, key).Result()
}
