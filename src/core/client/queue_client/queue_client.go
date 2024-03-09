package queueclient

import (
	"context"

	"github.com/snxl/stark_bank_integration/src/core/dto/webhook"
)

type QueueClient[FN any] interface {
	ProcessTask(fn ...map[string]FN) error
	IssueInvoiceDeliveryTask(ctx context.Context, obj interface{}) error
	ProcessInvoiceEvent(ctx context.Context, datum webhook.Event) error
}
