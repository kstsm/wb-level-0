package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kstsm/wb-level-0/consumer/config"
	"github.com/redis/go-redis/v9"
	"time"
)

type Redis struct {
	ctx    context.Context
	client *redis.Client
}

func NewRedis(ctx context.Context, cfg config.Config) *Redis {
	conf := &redis.Options{
		Addr: cfg.Redis.Address,
		DB:   cfg.Redis.DB,
	}

	rdb := redis.NewClient(conf)

	return &Redis{
		ctx:    ctx,
		client: rdb,
	}
}

func (r *Redis) SetJSON(key string, data interface{}, ttl time.Duration) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal value for cache: %w", err)
	}

	return r.client.Set(r.ctx, key, bytes, ttl).Err()
}

func (r *Redis) GetJSON(key string, dest interface{}) error {
	val, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		if errors.Is(redis.Nil, err) {
			return fmt.Errorf("ключ не найден")
		}
		return err
	}

	return json.Unmarshal([]byte(val), dest)
}

func (r *Redis) Close() error {
	return r.client.Close()
}
