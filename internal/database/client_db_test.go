package database_test

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"github.com/thgbianeck/bnck-ms-wallet/internal/database"
	"github.com/thgbianeck/bnck-ms-wallet/internal/entity"
	"testing"
)

type ClientDBTestSuit struct {
	suite.Suite
	db       *sql.DB
	clientDB *database.ClientDB
}

func (s *ClientDBTestSuit) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date, updated_at date)")
	s.clientDB = database.NewClientDB(db)
}

func (s *ClientDBTestSuit) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuit))
}

func (s *ClientDBTestSuit) TestSave() {
	client, _ := entity.NewClient("John Doe", "johndoe@mail.com")
	err := s.clientDB.Save(client)
	s.Nil(err)
}

func (s *ClientDBTestSuit) TestGet() {
	client, _ := entity.NewClient("John Doe", "johndoe@mail.com")
	s.clientDB.Save(client)

	clientDB, err := s.clientDB.Get(client.ID)
	s.Nil(err)
	s.Equal(client.ID, clientDB.ID)
	s.Equal(client.Name, clientDB.Name)
	s.Equal(client.Email, clientDB.Email)
}
