package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "johndoe@mail.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "johndoe@mail.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "johndoe@mail.com")

	err := client.Update("Jane Doe", "janedoe@mail.com")

	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Jane Doe", client.Name)
	assert.Equal(t, "janedoe@mail.com", client.Email)
	assert.True(t, client.UpdatedAt.After(client.CreatedAt))
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("John Doe", "johndoe@mail.com")
	err := client.Update("", "janedoe@mail.com")

	assert.NotNil(t, err)
	assert.Error(t, err, "name is required")
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("john doe", "j@j.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}
