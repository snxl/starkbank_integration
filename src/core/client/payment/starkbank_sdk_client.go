package payment

import (
	"fmt"

	"github.com/snxl/stark_bank_integration/src/config/keys"
	"github.com/snxl/stark_bank_integration/src/core/dto/payment"
	"github.com/starkbank/sdk-go/starkbank"
	"github.com/starkbank/sdk-go/starkbank/invoice"
	"github.com/starkbank/sdk-go/starkbank/transfer"
	"github.com/starkinfra/core-go/starkcore/user/project"
	"github.com/starkinfra/core-go/starkcore/utils/checks"
)

func init() {
	key := keys.GetKeys()
	starkbank.User = project.Project{
		Id:          key.ProjectId,
		PrivateKey:  checks.CheckPrivateKey(key.PrivateKey),
		Environment: checks.CheckEnvironment(key.Environment),
	}
}

type StarkbankSDKClient struct{}

func NewStarkbankSDKClient() *StarkbankSDKClient {
	return &StarkbankSDKClient{}
}

func (s *StarkbankSDKClient) IssueInvoice(input payment.IssueInvoiceDTO) error {
	_, err := invoice.Create(
		[]invoice.Invoice{
			{
				TaxId:  input.Cpf,
				Name:   input.Name,
				Amount: input.Amount,
				Due:    &input.Due,
				Tags:   []string{"imediate"},
			},
		}, nil,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			return fmt.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	return nil
}

func (s *StarkbankSDKClient) SendTransfer(input payment.SendTransferDTO) error {
	_, err := transfer.Create(
		[]transfer.Transfer{
			{
				Amount:        input.Amount,
				Name:          input.Name,
				TaxId:         input.TaxId,
				BankCode:      input.BankCode,
				BranchCode:    input.BranchCode,
				AccountNumber: input.AccountNumber,
				AccountType:   input.AccountType,
			},
		}, nil,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			return fmt.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	return nil
}
