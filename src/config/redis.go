package config

import (
	"context"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/snxl/stark_bank_integration/src/config/keys"
)

var client *redis.Client
var redisOnce sync.Once

func RedisConn() *redis.Client {
	redisOnce.Do(func() {
		key := keys.GetKeys()

		client = redis.NewClient(&redis.Options{
			Addr:     key.RedisURL,
			Password: key.RedisPassword,
		})

		_, err := client.Ping(context.TODO()).Result()
		if err != nil {
			panic(err)
		}
	})

	return client
}
