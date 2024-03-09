package processinvoiceevent

import (
	"fmt"
	"time"

	"github.com/snxl/stark_bank_integration/src/core/client/payment"
	dto "github.com/snxl/stark_bank_integration/src/core/dto/payment"
	"github.com/snxl/stark_bank_integration/src/core/repository/cache"
)

type ProcessInvoiceEventUsecase struct {
	paymentClient   payment.PaymentClient
	cacheRepository cache.CacheRepository
}

func NewProcessInvoiceEventUsecase(
	paymentClient payment.PaymentClient,
	cacheRepository cache.CacheRepository,
) *ProcessInvoiceEventUsecase {
	return &ProcessInvoiceEventUsecase{
		paymentClient:   paymentClient,
		cacheRepository: cacheRepository,
	}
}

func (usecase *ProcessInvoiceEventUsecase) Run(input Input) error {
	if input.Subscription == "invoice" {
		if input.Status == "paid" {
			key := fmt.Sprintf("invoice-%s-amount-%d-paid-event", input.Id, input.Amount)
			has, err := usecase.cacheRepository.Get(key)
			if err != nil {
				return err
			}
			if has != "" {
				return fmt.Errorf("invoice already processed")
			}

			err = usecase.paidInvoice(input.Amount)
			if err != nil {
				return err
			}

			err = usecase.cacheRepository.Set(key, true, time.Minute*5)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (usecase *ProcessInvoiceEventUsecase) paidInvoice(amount int) error {
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
