package storage

import (
	"OLO-backend/olo_service/internal/domain/models"
	"fmt"
)

type SaveWidgetStorage interface {
	GetWidget(widgetId int64) (*models.Widget, error)
	SaveWidget(widgetId int64, description string) error
}

// region for mysql Provider

func (s *InMysqlStorage) initTableWidgets() {
	db := s.mysqlProvider.DB
	// Создание таблицы auth_data, если она еще не существует
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS " + TableNameWidgets + " (" +
		"widgetId INT NOT NULL PRIMARY KEY, " +
		"description VARCHAR(255) NOT NULL UNIQUE, " +
		")")
	if err != nil {
		s.log.Error("Error creating "+TableNameWidgets+" table: ", err)
	}
}

func (s *InMysqlStorage) initTestDataForWidgets() {
	s.addAppWithIndex(&models.Widget{
		ID:          1,
		Description: "WaterTimer",
	})
}

func (s *InMysqlStorage) addAppWithIndex(widget *models.Widget) error {
	driver, err := s.mysqlProvider.Driver()
	if err != nil {
		return err
	}
	_, err = driver.NamedExec("INSERT INTO "+TableNameWidgets+" (`id`, `description`) VALUES (:id, :description)", widget)
	return err
}

func (s *InMysqlStorage) GetWidget(widgetId int64) (*models.Widget, error) {
	driver, err := s.mysqlProvider.Driver()
	if err != nil {
		s.log.Error("Error get driver", err)
		return nil, err
	}

	sub := &models.Widget{}
	rows, err := driver.NamedQuery(fmt.Sprintf("SELECT * FROM "+TableNameWidgets+" WHERE id = %d", widgetId), sub)
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

func (s *InMysqlStorage) SaveWidget(widgetId int64, description string) error {
	driver, err := s.mysqlProvider.Driver()
	if err != nil {
		s.log.Error("Error insert to database", err)
	}
	driver.NamedExec("INSERT INTO "+TableNameWidgets+" (`id`, `name`, `secret`) VALUES (:id, :name, :secret)", map[string]interface{}{
		"id":          widgetId,
		"description": description,
	})
	return err
}

// endregion
