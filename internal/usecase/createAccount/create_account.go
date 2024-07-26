package create_account

import (
	"github.com/thgbianeck/bnck-ms-wallet/internal/entity"
	"github.com/thgbianeck/bnck-ms-wallet/internal/gateway"
)

type CreateAccountDTO struct {
	ClientID string `json:"clientId"`
}

type CreateAccountOutputDTO struct {
	ID string
}

type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

func NewCreateAccountUseCase(a gateway.AccountGateway, c gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{AccountGateway: a, ClientGateway: c}
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountDTO) (*CreateAccountOutputDTO, error) {
	client, err := uc.ClientGateway.Get(input.ClientID)
	if err != nil {
		return nil, err
	}

	account := entity.NewAccount(client)
	err = uc.AccountGateway.Save(account)
	if err != nil {
		return nil, err
	}

	return &CreateAccountOutputDTO{ID: account.ID}, nil
}
