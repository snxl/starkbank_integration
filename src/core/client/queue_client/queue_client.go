package queueclient

import (
	"context"

	"github.com/hibiken/asynq"
)

type QueueClient interface {
	ProcessTask(fn ...map[string]func(context.Context, *asynq.Task) error) error
	IssueInvoiceDeliveryTask(ctx context.Context, obj interface{}) error
}
