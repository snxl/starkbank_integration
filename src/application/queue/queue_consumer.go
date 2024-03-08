package queue

import (
	"context"
	"log"

	"github.com/hibiken/asynq"
	issueinvoice "github.com/snxl/stark_bank_integration/src/application/handler/issue_invoice"
	queueclient "github.com/snxl/stark_bank_integration/src/core/client/queue_client"
	"github.com/snxl/stark_bank_integration/src/shared/constant"
)

type QueueConsumer struct {
	queue queueclient.QueueClient
}

func NewQueueConsumer(
	queueClient queueclient.QueueClient,
) *QueueConsumer {
	return &QueueConsumer{
		queue: queueClient,
	}
}

func (q *QueueConsumer) Start() {
	err := q.queue.ProcessTask(
		map[string]func(context.Context, *asynq.Task) error{
			constant.TaskIssueInvoice: issueinvoice.NewIssueInvoiceHandler().Run,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
