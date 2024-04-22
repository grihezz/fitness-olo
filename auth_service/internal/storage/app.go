// Package storage provides storage implementations for application data.
package storage

import (
	"OLO-backend/auth_service/internal/domain/models"
	"fmt"
)

// AppStorage defines methods for interacting with application data.
type AppStorage interface {
	GetApp(appID int) (*models.App, error)
	SaveApp(appID int, name string, secret string) error
}

// initTableApps initializes the apps table in MySQL storage.
func (s *InMysqlStorage) initTableApps() {
	db := s.mysqlProvider.DB
	// Create the apps table if it doesn't exist
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS " + TableNameApp + " (" +
		"id INT NOT NULL PRIMARY KEY, " +
		"name VARCHAR(20) NOT NULL UNIQUE, " +
		"secret VARCHAR(10) NOT NULL UNIQUE" +
		")")
	if err != nil {
		s.log.Error("Error creating "+TableNameApp+" table: ", err)
	}
}

// initTestDataForApps initializes test data for apps in MySQL storage.
func (s *InMysqlStorage) initTestDataForApps() {
	s.addAppWithIndex(&models.App{
		ID:     1,
		Name:   "Auth",
		Secret: "test-test",
	})
}

// addAppWithIndex adds an app with the specified index to MySQL storage.
func (s *InMysqlStorage) addAppWithIndex(app *models.App) error {
	driver, err := s.mysqlProvider.Driver()
	if err != nil {
		return err
	}
	_, err = driver.NamedExec("INSERT INTO "+TableNameApp+" (`id`, `name`, `secret`) VALUES (:id, :name, :secret)", app)
	return err
}

// GetApp retrieves an app by ID from MySQL storage.
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

// SaveApp saves an app to MySQL storage.
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
