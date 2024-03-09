package cache

import "time"

type CacheRepository interface {
	Get(key string) (data string, err error)
	Set(key string, data interface{}, expiration time.Duration) (err error)
	Del(key string) (err error)
}
