package storage

import (
	"auth/internal/domain/models"
	"auth/internal/storage/provider"
	"fmt"
	"log/slog"
)

const TableName = "auth_data"

type UserStorage interface {
	GetUser(email string) (*models.User, error)
	SaveUser(email string, passhash []byte) error
}

type InMysqlStorage struct {
	mysqlProvider *provider.MySQLProvider
	log           *slog.Logger
}

func NewInAuthMysqlStorage(log *slog.Logger, address, username, password, database string, port uint16) *InMysqlStorage {
	mySQLProvider, err := provider.NewMySQLProvider(address, port, username, password, database)
	if err != nil {
		panic(err)
	}

	result := &InMysqlStorage{
		mysqlProvider: mySQLProvider,
		log:           log,
	}
	result.initTables()
	//result.initRegsTestData()
	return result
}

func (s *InMysqlStorage) initTables() {
	db := s.mysqlProvider.DB
	// Создание таблицы auth_data, если она еще не существует
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS " + TableName + " (" +
		"id BIGINT NOT NULL AUTO_INCREMENT, " +
		"email VARCHAR(255) NOT NULL UNIQUE, " +
		"role VARCHAR(10) NOT NULL DEFAULT \"USER\", " +
		"password_hash VARCHAR(64) NOT NULL, " +
		"date_register TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, " +
		"PRIMARY KEY (id)" +
		")")
	if err != nil {
		s.log.Error("Error creating auth_data table: ", err)
	}

}

func (s *InMysqlStorage) initRegsTestData() {
	db := s.mysqlProvider.DB
	_, err := db.Exec("INSERT INTO app (app_id, name, secret) VALUES (1, 'test', 'test-test')")
	if err != nil {
		s.log.Error("Error insert to database", err)
	}
}

func (s *InMysqlStorage) GetUser(email string) (*models.User, error) {
	driver, err := s.mysqlProvider.Driver()
	if err != nil {
		s.log.Error("Error get data from database", err)
		return nil, err
	}

	sub := &models.User{}
	rows, err := driver.NamedQuery(fmt.Sprintf("SELECT * FROM "+TableName+" WHERE email = '%s'", email), sub)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, fmt.Errorf("rows not found")
	}
	err = rows.StructScan(&sub)
	return sub, err
}

func (s *InMysqlStorage) SaveUser(email string, passhash []byte) error {
	driver, err := s.mysqlProvider.Driver()
	if err != nil {
		s.log.Error("Error insert to database", err)
	}
	driver.NamedExec("INSERT INTO "+TableName+" (`email`, `password_hash`) VALUES (:email, :password_hash)", map[string]interface{}{
		"email":         email,
		"password_hash": passhash,
	})
	return err
}
