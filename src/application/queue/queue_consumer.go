package queue

import (
	"log"

	"github.com/hibiken/asynq"
	issueinvoice "github.com/snxl/stark_bank_integration/src/application/handler/issue_invoice"
	processinvoiceevent "github.com/snxl/stark_bank_integration/src/application/handler/process_invoice_event"
	queueclient "github.com/snxl/stark_bank_integration/src/core/client/queue_client"
	"github.com/snxl/stark_bank_integration/src/shared/constant"
)

type QueueConsumer struct {
	queue queueclient.QueueClient[asynq.HandlerFunc]
}

func NewQueueConsumer(
	queueClient queueclient.QueueClient[asynq.HandlerFunc],
) *QueueConsumer {
	return &QueueConsumer{
		queue: queueClient,
	}
}

func (q *QueueConsumer) Start() {
	err := q.queue.ProcessTask(
		map[string]asynq.HandlerFunc{
			constant.TaskIssueInvoice:        issueinvoice.NewIssueInvoiceHandler().Run,
			constant.TaskProcessInvoiceEvent: processinvoiceevent.NewProcessInvoiceEventHandler().Run,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
