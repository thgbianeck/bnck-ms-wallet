package createtransaction_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/thgbianeck/bnck-ms-wallet/internal/entity"
	createtransaction "github.com/thgbianeck/bnck-ms-wallet/internal/usecase/createTransaction"
	"github.com/thgbianeck/bnck-ms-wallet/internal/usecase/mocks"
	"testing"
)

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client1, _ := entity.NewClient("client1", "client1@com")
	account1 := entity.NewAccount(client1)
	account1.Credit(1000)

	client2, _ := entity.NewClient("client2", "client2@com")
	account2 := entity.NewAccount(client2)
	account2.Credit(1000)

	mockAccount := &mocks.AccountGatewayMock{}
	mockAccount.On("FindByID", account1.ID).Return(account1, nil)
	mockAccount.On("FindByID", account2.ID).Return(account2, nil)

	mockTransaction := &mocks.TransactionGatewayMock{}
	mockTransaction.On("Create", mock.Anything).Return(nil)

	inputDto := createtransaction.CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        500,
	}

	uc := createtransaction.NewCreateTransactionUseCase(mockTransaction, mockAccount)
	output, err := uc.Execute(inputDto)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockAccount.AssertExpectations(t)
	mockTransaction.AssertExpectations(t)
	mockAccount.AssertNumberOfCalls(t, "FindByID", 2)
	mockTransaction.AssertNumberOfCalls(t, "Create", 1)

}
