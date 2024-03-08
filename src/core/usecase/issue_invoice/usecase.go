package issueinvoice

import (
	"math/rand"
	"time"

	cpffaker "github.com/snxl/stark_bank_integration/src/core/client/cpf_faker"
	"github.com/snxl/stark_bank_integration/src/core/client/payment"
	dto "github.com/snxl/stark_bank_integration/src/core/dto/payment"
	"github.com/snxl/stark_bank_integration/src/shared/constant"
)

type IssueInvoiceUsecase struct {
	paymentClient payment.PaymentClient
	cpfClient     cpffaker.CPFFakerClient
}

func NewIssueInvoiceUsecase(
	paymentClient payment.PaymentClient,
	cpfClient cpffaker.CPFFakerClient,
) *IssueInvoiceUsecase {
	return &IssueInvoiceUsecase{
		paymentClient: paymentClient,
		cpfClient:     cpfClient,
	}
}

func (usecase *IssueInvoiceUsecase) Run(input Input) {
	due := time.Now().Add(time.Hour * 24)
	cpf := usecase.cpfClient.Generate()
	randomAmount := rand.Intn(constant.MaxInvoiceAmount)

	err := usecase.paymentClient.IssueInvoice(dto.IssueInvoiceDTO{
		Cpf:    cpf,
		Name:   input.Id,
		Amount: randomAmount,
		Due:    due,
	})
	if err != nil {
		panic(err)
	}
}
