package main

import (
	"github.com/snxl/stark_bank_integration/src/application/job"
	"github.com/snxl/stark_bank_integration/src/application/queue"
	"github.com/snxl/stark_bank_integration/src/application/server"
	"github.com/snxl/stark_bank_integration/src/config"
	queueclient "github.com/snxl/stark_bank_integration/src/core/client/queue_client"
)

func main() {
	asynqConf := config.GetAsynq()
	asynqueue := queueclient.NewQueueAsynq(*asynqConf.RedisOpt)

	queue.NewQueueConsumer(asynqueue).Start()

	err := job.NewJob(asynqueue).Start()
	if err != nil {
		panic(err)
	}

	err = server.NewServer().Start("8081")
	if err != nil {
		panic(err)
	}
}
