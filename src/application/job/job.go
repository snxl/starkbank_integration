package job

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
	uuid "github.com/satori/go.uuid"

	"github.com/go-co-op/gocron"
	queueclient "github.com/snxl/stark_bank_integration/src/core/client/queue_client"
)

type Job interface {
	Start() (err error)
	issueInvoices()
}

type JobImpl struct {
	queue queueclient.QueueClient[asynq.HandlerFunc]
}

func NewJob(
	queue queueclient.QueueClient[asynq.HandlerFunc],
) Job {
	return &JobImpl{
		queue: queue,
	}
}

func (job *JobImpl) Start() (err error) {
	s := gocron.NewScheduler(time.Local)

	_, err = s.Every(30).Second().Do(job.issueInvoices)
	if err != nil {
		return err
	}

	s.StartAsync()
	log.Println("Init jobs")
	return err
}

func (job *JobImpl) issueInvoices() {
	// max, min := 12, 8
	// randomInvoices := rand.Intn(max-min+1) + min

	for i := 1; i <= 1; i++ {
		err := job.queue.IssueInvoiceDeliveryTask(
			context.Background(),
			uuid.NewV4(),
		)
		if err != nil {
			fmt.Printf("[ISSUE INVOICE] error to delivery invoice request: %v", err)
		}
	}
}
