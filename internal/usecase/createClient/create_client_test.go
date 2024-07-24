package createclient_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	createclient "github.com/thgbianeck/bnck-ms-wallet/internal/usecase/createClient"
	"github.com/thgbianeck/bnck-ms-wallet/internal/usecase/mocks"
	"testing"
)

func TestCreateClientUseCase_Execute(t *testing.T) {
	m := &mocks.ClientGatewayMock{}
	m.On("Save", mock.Anything).Return(nil)

	uc := createclient.NewCreateClientUseCase(m)
	output, err := uc.Execute(createclient.CreateClientDTO{Name: "Richard", Email: "rich@email.com"})
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, "Richard", output.Name)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}
