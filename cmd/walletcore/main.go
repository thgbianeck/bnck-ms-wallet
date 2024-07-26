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
	"github.com/thgbianeck/bnck-ms-wallet/internal/web"
	"github.com/thgbianeck/bnck-ms-wallet/internal/web/webserver"
	"github.com/thgbianeck/bnck-ms-wallet/pkg/events"
	"log"
)

func main() {
	serverMysql := "localhost"
	serverAppPort := ":3000"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", serverMysql, "3306", "wallet")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Verifica se a conexão está realmente funcionando
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	log.Println("Connected to the database successfully")

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()
	//eventDispatcher.Register("TransactionCreated", handler)

	clientDb := database.NewClientDB(db)
	transactionDb := database.NewTransactionDB(db)
	accountDb := database.NewAccountDB(db)

	createClientUseCase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUseCase := create_account.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(transactionDb, accountDb, eventDispatcher, transactionCreatedEvent)

	webServer := webserver.NewWebServer(serverAppPort)

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webServer.AddHandler("/clients", clientHandler.CreateClient)
	webServer.AddHandler("/accounts", accountHandler.CreateAccount)
	webServer.AddHandler("/transactions", transactionHandler.CreateTransaction)

	// Início do servidor web
	log.Println("Starting web server in port " + serverAppPort)
	err = webServer.Start()
	if err != nil {
		log.Fatalf("Error starting web server: %v", err)
	}
	log.Println("Web server started successfully")
}
