// Package storage provides storage implementations for various data entities.
package storage

import (
	"OLO-backend/auth_service/internal/storage/provider"
	"errors"
	"log/slog"
)

// Custom errors for user and app operations.
var (
	ErrUserExist    = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
	ErrAppNotFound  = errors.New("app not found")

	// Table names in the database.
	TableNameUser = "users"
	TableNameApp  = "app_table"
)

// InMysqlStorage represents the MySQL storage implementation.
type InMysqlStorage struct {
	mysqlProvider *provider.MySQLProvider
	log           *slog.Logger
}

// NewInAuthMysqlStorage creates a new instance of InMysqlStorage.
func NewInAuthMysqlStorage(log *slog.Logger, address, username, password, database string, port uint16) *InMysqlStorage {
	mySQLProvider, err := provider.NewMySQLProvider(address, port, username, password, database)
	if err != nil {
		panic(err)
	}

	result := &InMysqlStorage{
		mysqlProvider: mySQLProvider,
		log:           log,
	}
	result.init()
	return result
}

// init initializes the MySQL storage.
func (s *InMysqlStorage) init() {
	s.initTableUser()

	s.initTableApps()
	s.initTestDataForApps()
}
