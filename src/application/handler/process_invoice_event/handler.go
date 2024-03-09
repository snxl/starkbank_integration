package processinvoiceevent

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
	"github.com/snxl/stark_bank_integration/src/core/client/payment"
	"github.com/snxl/stark_bank_integration/src/core/dto/webhook"
	"github.com/snxl/stark_bank_integration/src/core/repository/cache"
	processinvoiceevent "github.com/snxl/stark_bank_integration/src/core/usecase/process_invoice_event"
)

type ProcessInvoiceEventHandler struct{}

func NewProcessInvoiceEventHandler() *ProcessInvoiceEventHandler {
	return &ProcessInvoiceEventHandler{}
}

func (i *ProcessInvoiceEventHandler) Run(ctx context.Context, task *asynq.Task) error {
	var event webhook.Event
	err := json.Unmarshal(task.Payload(), &event)
	if err != nil {
		return err
	}

	usecase := processinvoiceevent.NewProcessInvoiceEventUsecase(
		payment.NewStarkbankSDKClient(),
		cache.NewRedisRepository(),
	)
	err = usecase.Run(processinvoiceevent.Input{
		Subscription: event.Subscription,
		Status:       event.Log.Invoice.Status,
		Amount:       event.Log.Invoice.Amount,
		Id:           event.Log.Invoice.Id,
	})
	if err != nil {
		return err
	}

	return nil
}
