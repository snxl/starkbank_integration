package payment

import "github.com/snxl/stark_bank_integration/src/core/dto/payment"

type PaymentClient interface {
	IssueInvoice(input payment.IssueInvoiceDTO) error
	SendTransfer(input payment.SendTransferDTO) error
}
