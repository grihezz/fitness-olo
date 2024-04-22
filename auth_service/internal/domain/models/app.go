// Package models provides data models for the application.
package models

// App represents a data structure for application entities.
type App struct {
	ID     int    `db:"id"`
	Name   string `db:"name"`
	Secret string `db:"secret"`
}
