package storage

import (
	"OLO-backend/auth_service/internal/domain/models"
	"fmt"
)

type AppStorage interface {
	GetApp(appID int) (*models.App, error)
	SaveApp(appID int, name string, secret string) error
}

// region for mysql Provider

func (s *InMysqlStorage) initTableApps() {
	db := s.mysqlProvider.DB
	// Создание таблицы auth_data, если она еще не существует
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS " + TableNameApp + " (" +
		"id INT NOT NULL PRIMARY KEY, " +
		"name VARCHAR(255) NOT NULL UNIQUE, " +
		"secret VARCHAR(10) NOT NULL UNIQUE" +
		")")
	if err != nil {
		s.log.Error("Error creating "+TableNameApp+" table: ", err)
	}
}

func (s *InMysqlStorage) initTestDataForApps() {
	s.addAppWithIndex(&models.App{
		ID:     1,
		Name:   "Auth",
		Secret: "test-test",
	})
}

func (s *InMysqlStorage) addAppWithIndex(app *models.App) error {
	driver, err := s.mysqlProvider.Driver()
	if err != nil {
		return err
	}
	_, err = driver.NamedExec("INSERT INTO "+TableNameApp+" (`id`, `name`, `secret`) VALUES (:id, :name, :secret)", app)
	return err
}

func (s *InMysqlStorage) GetApp(appID int) (*models.App, error) {
	driver, err := s.mysqlProvider.Driver()
	if err != nil {
		s.log.Error("Error get driver", err)
		return nil, err
	}

	sub := &models.App{}
	rows, err := driver.NamedQuery(fmt.Sprintf("SELECT * FROM "+TableNameApp+" WHERE id = %d", appID), sub)
	if err != nil {
		s.log.Error("Error get data from database", err)
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, fmt.Errorf("rows not found")
	}
	err = rows.StructScan(&sub)
	return sub, err
}

func (s *InMysqlStorage) SaveApp(appID int, name string, secret string) error {
	driver, err := s.mysqlProvider.Driver()
	if err != nil {
		s.log.Error("Error insert to database", err)
	}
	driver.NamedExec("INSERT INTO "+TableNameApp+" (`id`, `name`, `secret`) VALUES (:id, :name, :secret)", map[string]interface{}{
		"id":     appID,
		"name":   name,
		"secret": secret,
	})
	return err
}

// endregion
