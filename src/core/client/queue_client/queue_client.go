package queueclient

import (
	"context"
)

type QueueClient[FN any] interface {
	ProcessTask(fn ...map[string]FN) error
	IssueInvoiceDeliveryTask(ctx context.Context, obj interface{}) error
}
