package createtransaction

import (
	"github.com/thgbianeck/bnck-ms-wallet/internal/entity"
	"github.com/thgbianeck/bnck-ms-wallet/internal/gateway"
)

type CreateTransactionInputDTO struct {
	AccountIDFrom string  `json:"account_id_from"`
	AccountIDTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionOutputDTO struct {
	ID            string  `json:"id"`
	AccountIDFrom string  `json:"account_id_from"`
	AccountIDTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionUseCase struct {
	TransactionGateway gateway.TransactionGateway
	AccountGateway     gateway.AccountGateway
}

func NewCreateTransactionUseCase(transactionGateway gateway.TransactionGateway, accountGateway gateway.AccountGateway) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{TransactionGateway: transactionGateway, AccountGateway: accountGateway}
}

func (uc *CreateTransactionUseCase) Execute(input CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	accountFrom, err := uc.AccountGateway.FindByID(input.AccountIDFrom)
	if err != nil {
		return nil, err
	}

	accountTo, err := uc.AccountGateway.FindByID(input.AccountIDTo)
	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	if err != nil {
		return nil, err
	}

	err = uc.TransactionGateway.Create(transaction)
	if err != nil {
		return nil, err
	}

	return &CreateTransactionOutputDTO{
		ID:            transaction.ID,
		AccountIDFrom: transaction.AccountFrom.ID,
		AccountIDTo:   transaction.AccountTo.ID,
		Amount:        transaction.Amount,
	}, nil
}
