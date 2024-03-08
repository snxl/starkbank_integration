package queueclient

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/snxl/stark_bank_integration/src/shared/constant"
)

type QueueAsynq struct {
	client *asynq.Client
	server *asynq.Server
}

func NewQueueAsynq(
	redisOpt asynq.RedisClientOpt,
) *QueueAsynq {
	client := asynq.NewClient(redisOpt)
	server := asynq.NewServer(redisOpt,
		asynq.Config{
			Concurrency: 4,
		})
	return &QueueAsynq{
		client: client,
		server: server,
	}
}

func (q *QueueAsynq) ProcessTask(fn ...map[string]asynq.HandlerFunc) error {
	mux := asynq.NewServeMux()

	for _, maps := range fn {
		for key, value := range maps {
			mux.HandleFunc(key, value)
		}
	}

	return q.server.Start(mux)
}

func (q *QueueAsynq) IssueInvoiceDeliveryTask(ctx context.Context, obj interface{}) error {
	jsonPayload, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %w", err)
	}

	task := asynq.NewTask(constant.TaskIssueInvoice, jsonPayload)
	info, err := q.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %w", err)
	}

	msg := fmt.Sprintf("type %s payload %s queue %s max retry %d", task.Type(), string(task.Payload()), info.Queue, info.MaxRetry)
	fmt.Println(msg)

	return nil
}
