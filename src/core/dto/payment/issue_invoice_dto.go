package payment

import "time"

type IssueInvoiceDTO struct {
	Name   string
	Cpf    string
	Amount int
	Due    time.Time
}
