package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/snxl/stark_bank_integration/src/config"
)

type RedisRepository struct {
	redis *redis.Client
}

func NewRedisRepository() *RedisRepository {
	return &RedisRepository{
		redis: config.RedisConn(),
	}
}

func (r *RedisRepository) Set(key string, data interface{}, expiration time.Duration) (err error) {
	err = r.redis.Set(context.Background(), key, data, expiration).Err()
	return
}

func (r *RedisRepository) Get(key string) (data string, err error) {
	data, err = r.redis.Get(context.Background(), key).Result()
	if err == redis.Nil {
		err = nil
	}
	return
}

func (r *RedisRepository) Del(key string) (err error) {
	err = r.redis.Del(context.Background(), key).Err()
	return
}
