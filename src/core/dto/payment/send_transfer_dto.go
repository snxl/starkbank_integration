package payment

type SendTransferDTO struct {
	BankCode      string
	BranchCode    string
	AccountNumber string
	TaxId         string
	Name          string
	AccountType   string
	Amount        int
}
