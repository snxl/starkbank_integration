package keys

import (
	"os"

	"github.com/joho/godotenv"
)

type Envs struct {
	RedisURL      string
	RedisPassword string
	ProjectId     string
	PrivateKey    string
	Environment   string
}

func init() {
	_ = godotenv.Load()
}

func GetKeys() Envs {
	return Envs{
		RedisURL:      os.Getenv("REDIS_URL"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		ProjectId:     os.Getenv("STARK_BANK_PROJECT_ID"),
		PrivateKey:    os.Getenv("STARK_BANK_PRIVATE_KEY"),
		Environment:   os.Getenv("STARK_BANK_ENVIRONMENT"),
	}
}
