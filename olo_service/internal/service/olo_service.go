package service

import (
	"OLO-backend/olo_service/internal/entity"
	"OLO-backend/olo_service/internal/repository"
	"fmt"
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
	widgets, err := s.repo.GetAllWidgets()
	if err != nil {
		return nil, fmt.Errorf("can't get all widgets")
	}
	return widgets, nil
}

func (s *OloService) GetUserWidgets(userId int64) ([]entity.Widget, error) {
	widgets, err := s.repo.GetUserWidgets(userId)
	if err != nil {
		return nil, fmt.Errorf("can't get widgets of user")
	}
	return widgets, nil
}

func (s *OloService) AddWidgetForUser(widgetId, userId int64) error {
	err := s.repo.AddWidgetForUser(widgetId, userId)
	if err != nil {
		return fmt.Errorf("can't add widget for user")
	}
	return nil
}

func (s *OloService) GetAllArticles() ([]entity.Article, error) {
	articles, err := s.repo.GetAllArticles()
	if err != nil {
		return nil, fmt.Errorf("can't get all articles")
	}
	return articles, nil
}

func (s *OloService) GetUsersArticles(userId int64) ([]entity.Article, error) {
	articles, err := s.repo.GetUsersArticles(userId)
	if err != nil {
		return nil, fmt.Errorf("can't get articles of user")
	}
	return articles, nil
}

func (s *OloService) AddArticleForUser(articleId, userId int64) error {
	err := s.repo.AddArticleForUser(articleId, userId)
	if err != nil {
		return fmt.Errorf("can't add article for user")
	}
	return nil
}
