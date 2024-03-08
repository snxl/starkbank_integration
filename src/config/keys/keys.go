package keys

import (
	"os"

	"github.com/joho/godotenv"
)

type Envs struct {
	RedisURL      string
	RedisPassword string
}

func init() {
	_ = godotenv.Load()
}

func GetKeys() Envs {
	return Envs{
		RedisURL:      os.Getenv("REDIS_URL"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
	}
}
