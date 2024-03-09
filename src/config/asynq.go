package config

import (
	"sync"

	"github.com/hibiken/asynq"
	"github.com/snxl/stark_bank_integration/src/config/keys"
)

type AsynqConf struct {
	RedisOpt *asynq.RedisClientOpt
	Client   asynq.Client
	Server   asynq.Server
}

var asynqOnce sync.Once
var instance *AsynqConf

func GetAsynq() *AsynqConf {
	key := keys.GetKeys()

	asynqOnce.Do(func() {
		instance = &AsynqConf{}
		instance.RedisOpt = &asynq.RedisClientOpt{Addr: key.RedisURL, Password: key.RedisPassword}
		instance.Client = *asynq.NewClient(instance.RedisOpt)
		instance.Server = *asynq.NewServer(instance.RedisOpt, asynq.Config{
			Concurrency: 4,
		})
	})

	return instance
}
