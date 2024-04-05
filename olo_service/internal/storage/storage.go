package storage

import (
	"OLO-backend/olo_service/internal/storage/provider"
	"errors"
	"log/slog"
)

var (
	ErrUserExist    = errors.New("user already exist")
	ErrUserNotFound = errors.New("user not found")
	ErrAppNotFound  = errors.New("app not found")

	TableNameUser    = "users"
	TableNameWidgets = "wigets"
)

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
	result.init()
	return result
}

func (s *InMysqlStorage) init() {
	s.initTableWidgets()
	s.initTestDataForWidgets()
}
