package create_client

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/thgbianeck/bnck-ms-wallet/internal/usecase/mocks"
	"testing"
)

func TestCreateClientUseCase_Execute(t *testing.T) {
	m := &mocks.ClientGatewayMock{}
	m.On("Save", mock.Anything).Return(nil)

	uc := NewCreateClientUseCase(m)
	output, err := uc.Execute(CreateClientInputDTO{Name: "Richard", Email: "rich@email.com"})
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, "Richard", output.Name)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}
