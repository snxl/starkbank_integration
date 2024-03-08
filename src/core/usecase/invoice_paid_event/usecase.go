package invoicepaidevent

import (
	"fmt"

	"github.com/snxl/stark_bank_integration/src/core/client/payment"
	dto "github.com/snxl/stark_bank_integration/src/core/dto/payment"
)

type InvoicePaidEventUsecase struct {
	paymentClient payment.PaymentClient
}

func NewInvoicePaidEventUsecase(
	paymentClient payment.PaymentClient,
) *InvoicePaidEventUsecase {
	return &InvoicePaidEventUsecase{
		paymentClient: paymentClient,
	}
}

func (usecase *InvoicePaidEventUsecase) Run(input Input) error {
	fmt.Println(input)
	if input.Subscription == "invoice" {
		if input.Status == "paid" {
			return usecase.paidInvoice(input.Amount)
		}
	}
	return nil
}

func (usecase *InvoicePaidEventUsecase) paidInvoice(amount int) error {
	err := usecase.paymentClient.SendTransfer(dto.SendTransferDTO{
		Amount:        amount,
		BankCode:      "20018183",
		BranchCode:    "0001",
		AccountNumber: "6341320293482496",
		Name:          "Stark Bank S.A.",
		TaxId:         "20.018.183/0001-80",
		AccountType:   "payment",
	})
	if err != nil {
		return err
	}
	return nil
}
