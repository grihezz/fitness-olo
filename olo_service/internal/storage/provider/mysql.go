package provider

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySQLProvider struct {
	Username string
	Password string
	Database string
	Address  string
	Port     uint16

	DB *sqlx.DB
}

func NewMySQLProvider(address string, port uint16, username, password, database string) (*MySQLProvider, error) {
	provider := &MySQLProvider{
		Username: username,
		Password: password,
		Database: database,
		Address:  address,
		Port:     port,
	}

	err := provider.init()

	if err != nil {
		return nil, err
	}

	return provider, nil
}

func (provider *MySQLProvider) init() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", provider.Username, provider.Password, provider.Address, provider.Port, provider.Database)
	driver, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}

	driver.SetMaxOpenConns(3)
	driver.SetMaxIdleConns(2)

	provider.DB = driver
	return nil
}

func (provider *MySQLProvider) Driver() (*sqlx.DB, error) {
	driver := provider.DB
	if driver == nil {
		return nil, errors.New("database can't connect to mysql")
	}
	return driver, nil
}
