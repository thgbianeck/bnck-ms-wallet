package create_account

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/thgbianeck/bnck-ms-wallet/internal/entity"
	"github.com/thgbianeck/bnck-ms-wallet/internal/usecase/mocks"
	"testing"
)

func TestCreateAccountUseCase_Execute(t *testing.T) {
	client, _ := entity.NewClient("richard", "rich@mail.com")
	clientMock := &mocks.ClientGatewayMock{}
	clientMock.On("Get", client.ID).Return(client, nil)

	accountMock := &mocks.AccountGatewayMock{}
	accountMock.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUseCase(accountMock, clientMock)
	inputDto := CreateAccountInputDTO{ClientID: client.ID}
	output, err := uc.Execute(inputDto)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	clientMock.AssertExpectations(t)
	accountMock.AssertExpectations(t)
	clientMock.AssertNumberOfCalls(t, "Get", 1)
	accountMock.AssertNumberOfCalls(t, "Save", 1)

}
