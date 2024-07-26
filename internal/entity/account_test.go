package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "john.doe@mail.com")
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestCreateAccountWithoutClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "john.doe@mail.com")
	account := NewAccount(client)
	account.Credit(1000)
	assert.Equal(t, float64(1000), account.Balance)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "john.doe@mail.com")
	account := NewAccount(client)
	account.Credit(1000)
	account.Debit(500)
	assert.Equal(t, float64(500), account.Balance)
}
