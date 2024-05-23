// Package repository provides interfaces and implementations for interacting with data entities.
//
// This package defines interfaces for the Widget and Article entities, specifying methods for various operations.
// It also includes a Repository struct that combines both interfaces to provide a unified interface for data access.
package repository

import (
	"OLO-backend/olo_service/internal/entity"
	"OLO-backend/olo_service/internal/repository/provider"
)

// Widget represents the interface for interacting with widget data.
type Widget interface {
	GetAllWidgets() ([]entity.Widget, error)
	AddWidgetForUser(widgetId, userId int64) error
	GetUserWidgets(userId int64) ([]entity.Widget, error)
	DeleteWidgetForUser(widgetId int64, userId int64) error
}

// Article represents the interface for interacting with article data.
type Article interface {
	AddArticleForUser(articleId, userId int64) error
	GetAllArticles() ([]entity.Article, error)
	GetUsersArticles(userid int64) ([]entity.Article, error)
	DeleteArticleForUser(articleId int64, userId int64) error
}

// Repository represents a unified interface for interacting with both widget and article data.
type Repository struct {
	Widget  // Widget interface for widget-related operations
	Article // Article interface for article-related operations
}

// NewRepository creates a new instance of Repository with the provided MySQLProvider.
func NewRepository(mysqlProvider *provider.MySQLProvider) *Repository {
	return &Repository{
		Widget:  NewWidgetRepo(mysqlProvider),
		Article: NewArticleRepo(mysqlProvider),
	}
}
