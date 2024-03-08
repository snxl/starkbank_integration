package main

import (
	"github.com/hibiken/asynq"
	"github.com/snxl/stark_bank_integration/src/application/job"
	"github.com/snxl/stark_bank_integration/src/application/queue"
	"github.com/snxl/stark_bank_integration/src/application/server"
	"github.com/snxl/stark_bank_integration/src/config/keys"
	queueclient "github.com/snxl/stark_bank_integration/src/core/client/queue_client"
)

func main() {
	key := keys.GetKeys()

	redisOpt := &asynq.RedisClientOpt{Addr: key.RedisURL, Password: key.RedisPassword}
	asynqueue := queueclient.NewQueueAsynq(*redisOpt)

	queue.NewQueueConsumer(asynqueue).Start()

	err := job.NewJob(asynqueue).Start()
	if err != nil {
		panic(err)
	}

	err = server.NewServer().Start("8080")
	if err != nil {
		panic(err)
	}
}
