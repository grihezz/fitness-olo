package service

import (
	"OLO-backend/olo_service/internal/entity"
	"OLO-backend/olo_service/internal/repository"
	"log/slog"
)

type OloService struct {
	log  *slog.Logger
	repo *repository.Repository
}

func NewOloService(log *slog.Logger, repo *repository.Repository) *OloService {
	return &OloService{
		repo: repo,
		log:  log,
	}
}

func (s *OloService) GetAllWidgets() ([]entity.Widget, error) {
	return s.repo.GetAllWidgets()
}

func (s *OloService) GetUserWidgets(userId int64) ([]entity.Widget, error) {
	return s.repo.GetUserWidgets(userId)
}

func (s *OloService) AddWidgetForUser(widgetId, userId int64) error {
	return s.repo.AddWidgetForUser(widgetId, userId)
}

func (s *OloService) GetAllArticles() ([]entity.Article, error) {
	return s.repo.GetAllArticles()
}

func (s *OloService) GetUsersArticles(userId int64) ([]entity.Article, error) {
	return s.repo.GetUsersArticles(userId)
}

func (s *OloService) AddArticleForUser(articleId, userId int64) error {
	return s.repo.AddArticleForUser(articleId, userId)
}
