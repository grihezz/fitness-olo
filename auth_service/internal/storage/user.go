// Package storage provides storage implementations for various data entities.
package storage

import (
	"OLO-backend/auth_service/internal/domain/models"
	"fmt"
)

// UserStorage defines methods for interacting with user data.
type UserStorage interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id int64) (*models.User, error)
	SaveUser(email string, passhash []byte) (int64, error)
}

// initTableUser initializes the users table in MySQL storage.
func (s *InMysqlStorage) initTableUser() {
	db := s.mysqlProvider.DB
	// Create the users table if it doesn't exist
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS " + TableNameUser + " (" +
		"id BIGINT NOT NULL AUTO_INCREMENT, " +
		"email VARCHAR(35) NOT NULL UNIQUE, " +
		"role VARCHAR(5) NOT NULL DEFAULT \"USER\", " +
		"password_hash VARCHAR(64) NOT NULL, " +
		"date_register TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, " +
		"PRIMARY KEY (id)" +
		")")
	if err != nil {
		s.log.Error("Error creating "+TableNameUser+" table: ", err)
	}
}

// GetUserByEmail retrieves a user by email from MySQL storage.
func (s *InMysqlStorage) GetUserByEmail(email string) (*models.User, error) {
	driver, err := s.mysqlProvider.Driver()
	if err != nil {
		s.log.Error("Error getting data from database", err)
		return nil, err
	}

	sub := &models.User{}
	rows, err := driver.NamedQuery(fmt.Sprintf("SELECT * FROM "+TableNameUser+" WHERE email = '%s'", email), sub)
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

// GetUserById retrieves a user by ID from MySQL storage.
func (s *InMysqlStorage) GetUserById(id int64) (*models.User, error) {
	driver, err := s.mysqlProvider.Driver()
	if err != nil {
		s.log.Error("Error getting data from database", err)
		return nil, err
	}

	sub := &models.User{}
	rows, err := driver.NamedQuery(fmt.Sprintf("SELECT * FROM "+TableNameUser+" WHERE id = '%d'", id), sub)
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

// SaveUser saves a user to MySQL storage.
func (s *InMysqlStorage) SaveUser(email string, passhash []byte) (int64, error) {
	driver, err := s.mysqlProvider.Driver()
	if err != nil {
		s.log.Error("Error inserting to database", err)
	}
	res, err := driver.NamedExec("INSERT INTO "+TableNameUser+" (`email`, `password_hash`) VALUES (:email, :password_hash)", map[string]interface{}{
		"email":         email,
		"password_hash": passhash,
	})
	if err != nil {
		s.log.Error("Error saving user", err)
		return 0, err
	}
	id, err := res.LastInsertId()
	return id, err
}
