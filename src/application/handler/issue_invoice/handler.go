package issueinvoice

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
	cpffaker "github.com/snxl/stark_bank_integration/src/core/client/cpf_faker"
	"github.com/snxl/stark_bank_integration/src/core/client/payment"
	issueinvoice "github.com/snxl/stark_bank_integration/src/core/usecase/issue_invoice"
)

type IssueInvoiceHandler struct{}

func NewIssueInvoiceHandler() *IssueInvoiceHandler {
	return &IssueInvoiceHandler{}
}

func (i *IssueInvoiceHandler) Run(ctx context.Context, task *asynq.Task) error {
	var invoiceId string
	err := json.Unmarshal(task.Payload(), &invoiceId)
	if err != nil {
		return err
	}

	usecase := issueinvoice.NewIssueInvoiceUsecase(
		payment.NewStarkbankSDKClient(),
		cpffaker.NewGoCPF(),
	)
	usecase.Run(issueinvoice.Input{Id: invoiceId})

	return nil
}
