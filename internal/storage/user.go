package storage

import (
	"auth/internal/storage/provider"
	"log/slog"
)

const TableName = "auth_data"

type AuthStorage interface {
	//RegUser(username, email, password, role string) error
	//ChangeRoleByUsername(username, role string) error
	//FindByUsername(username string) (*db.UserCredential, error)

	SaveUser(email string, passhash []byte) error
}

type InAuthMysqlStorage struct {
	mysqlProvider *provider.MySQLProvider

	log *slog.Logger
}

func NewInAuthMysqlRepository(log *slog.Logger, address, username, password, database string, port uint16) AuthStorage {
	mySQLProvider, err := provider.NewMySQLProvider(address, port, username, password, database)
	if err != nil {
		panic(err)
	}

	result := &InAuthMysqlStorage{
		mysqlProvider: mySQLProvider,
		log:           log,
	}
	result.initTables()
	//result.initRegsTestData()
	return result
}

func (s *InAuthMysqlStorage) initTables() {
	db := s.mysqlProvider.DB
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS " + TableName + " (" +
		"id BIGINT NOT NULL AUTO_INCREMENT, " +
		"email VARCHAR(255) NOT NULL UNIQUE, " +
		"role VARCHAR(10) NOT NULL DEFAULT \"USER\", " +
		"password_hash VARCHAR(64) NOT NULL, " +
		"date_register TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, " +
		"PRIMARY KEY (id)" +
		")")
	if err != nil {
		//utils.GetLogger().Fatal(err)
	}
}

func (s *InAuthMysqlStorage) SaveUser(email string, passhash []byte) error {
	panic("Not implements!")
}
