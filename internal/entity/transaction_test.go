package entity_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/thgbianeck/bnck-ms-wallet/internal/entity"
	"testing"
)

func TestCreateTransaction(t *testing.T) {
	client1, _ := entity.NewClient("rich 1", "rich@com")
	account1 := entity.NewAccount(client1)
	client2, _ := entity.NewClient("john doe", "j@j.com")
	account2 := entity.NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := entity.NewTransaction(account2, account1, 100)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, 1100.0, account1.Balance)
	assert.Equal(t, 900.0, account2.Balance)
}

func TestCreateTransactionWithoutBalance(t *testing.T) {
	client1, _ := entity.NewClient("rich 1", "rich@com")
	account1 := entity.NewAccount(client1)
	client2, _ := entity.NewClient("john doe", "j@j.com")
	account2 := entity.NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := entity.NewTransaction(account2, account1, 2000)
	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Equal(t, 1000.0, account1.Balance)
	assert.Equal(t, 1000.0, account2.Balance)
	assert.Error(t, err, "insufficient funds")
}
