package gateway

import "github.com/thgbianeck/bnck-ms-wallet/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
