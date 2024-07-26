package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/thgbianeck/bnck-ms-wallet/internal/database"
	"github.com/thgbianeck/bnck-ms-wallet/internal/event"
	create_account "github.com/thgbianeck/bnck-ms-wallet/internal/usecase/createAccount"
	create_client "github.com/thgbianeck/bnck-ms-wallet/internal/usecase/createClient"
	create_transaction "github.com/thgbianeck/bnck-ms-wallet/internal/usecase/createTransaction"
	"github.com/thgbianeck/bnck-ms-wallet/pkg/events"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()
	//eventDispatcher.Register("TransactionCreated", handler)

	clientDb := database.NewClientDB(db)
	transactionDb := database.NewTransactionDB(db)
	accountDb := database.NewAccountDB(db)

	createClientUseCase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUseCase := create_account.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(transactionDb, accountDb, eventDispatcher, transactionCreatedEvent)
}
