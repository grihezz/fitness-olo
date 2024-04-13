package repository

import (
	"OLO-backend/olo_service/internal/entity"
	"OLO-backend/olo_service/internal/repository/provider"
)

type Widget interface {
	GetAllWidgets() ([]entity.Widget, error)
	AddWidgetForUser(widgetId, userId int64) error
	GetUserWidgets(userId int64) ([]entity.Widget, error)
}

type Article interface {
	AddArticleForUser(articleId, userId int64) error
	GetAllArticles() ([]entity.Article, error)
	GetUsersArticles(userid int64) ([]entity.Article, error)
}

type Repository struct {
	Widget
	Article
}

func NewRepository(mysqlProvider *provider.MySQLProvider) *Repository {
	return &Repository{
		Widget:  NewWidgetRepo(mysqlProvider),
		Article: NewArticleRepo(mysqlProvider),
	}
}
