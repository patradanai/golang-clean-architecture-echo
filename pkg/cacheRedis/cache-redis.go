package cacheredis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type cacheRedis struct {
	rds *redis.Client
}

type cacheEntry struct {
}

type CacheRedis interface {
	BuildKey(group string, key string) string

	Get(ctx context.Context, key string) (interface{}, error)

	Set(ctx context.Context, key string, value interface{}) error

	Del(ctx context.Context, key string) error
}

func InitRedisCache(rds *redis.Client) CacheRedis {
	return &cacheRedis{
		rds,
	}
}

func (c *cacheRedis) BuildKey(group string, key string) string {
	return fmt.Sprintf("%v%v", group, key)
}

func (c *cacheRedis) Get(ctx context.Context, key string) (interface{}, error) {
	val, err := c.rds.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (c *cacheRedis) Set(ctx context.Context, key string, value interface{}) error {
	_, err := c.rds.Set(ctx, key, value, 0).Result()
	if err != nil {
		return err
	}

	return nil
}

func (c *cacheRedis) SetWithExpire(ctx context.Context, key string, value interface{}, expire time.Duration) error {
	_, err := c.rds.SetNX(ctx, key, value, expire).Result()
	if err != nil {
		return err
	}

	return nil
}

func (c *cacheRedis) Del(ctx context.Context, key string) error {
	_, err := c.rds.Del(ctx, key).Result()
	if err != nil {
		return err
	}

	return nil
}
