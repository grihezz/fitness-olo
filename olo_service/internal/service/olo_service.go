// Package service provides functions for interacting with various data entities.
//
// This package includes functions for the following operations:
//   - GetAllWidgets: Retrieve all widgets from the repository.
//   - GetUserWidgets: Retrieve widgets associated with a specific user from the repository.
//   - AddWidgetForUser: Add a widget for a specific user.
//   - GetAllArticles: Retrieve all articles from the repository.
//   - GetUsersArticles: Retrieve articles associated with a specific user from the repository.
//   - AddArticleForUser: Add an article for a specific user.
package service

import (
	"OLO-backend/olo_service/internal/entity"
	"OLO-backend/olo_service/internal/repository"
	"fmt"
	"log/slog"
)

// OloService represents the service for OLO operations.
type OloService struct {
	log  *slog.Logger           // Logging
	repo *repository.Repository // Repository for OLO
}

// NewOloService creates a new instance of OloService with the provided logger and repository.
func NewOloService(log *slog.Logger, repo *repository.Repository) *OloService {
	return &OloService{
		repo: repo,
		log:  log,
	}
}

// GetAllWidgets retrieves all widgets from the repository.
func (s *OloService) GetAllWidgets() ([]entity.Widget, error) {
	widgets, err := s.repo.GetAllWidgets()
	if err != nil {
		return nil, fmt.Errorf("can't get all widgets")
	}
	return widgets, nil
}

// GetUserWidgets retrieves widgets associated with a specific user from the repository.
func (s *OloService) GetUserWidgets(userId int64) ([]entity.Widget, error) {
	widgets, err := s.repo.GetUserWidgets(userId)
	if err != nil {
		return nil, fmt.Errorf("can't get widgets of user")
	}
	return widgets, nil
}

// AddWidgetForUser adds a widget for a specific user.
func (s *OloService) AddWidgetForUser(widgetId, userId int64) error {
	err := s.repo.AddWidgetForUser(widgetId, userId)
	if err != nil {
		return fmt.Errorf("can't add widget for user")
	}
	return nil
}

// GetAllArticles retrieves all articles from the repository.
func (s *OloService) GetAllArticles() ([]entity.Article, error) {
	articles, err := s.repo.GetAllArticles()
	if err != nil {
		return nil, fmt.Errorf("can't get all articles")
	}
	return articles, nil
}

// GetUsersArticles retrieves articles associated with a specific user from the repository.
func (s *OloService) GetUsersArticles(userId int64) ([]entity.Article, error) {
	articles, err := s.repo.GetUsersArticles(userId)
	if err != nil {
		return nil, fmt.Errorf("can't get articles of user")
	}
	return articles, nil
}

// AddArticleForUser adds an article for a specific user.
func (s *OloService) AddArticleForUser(articleId, userId int64) error {
	err := s.repo.AddArticleForUser(articleId, userId)
	if err != nil {
		return fmt.Errorf("can't add article for user")
	}
	return nil
}

func (s *OloService) DeleteArticleForUser(articleId int64, userId int64) error {
	err := s.repo.DeleteArticleForUser(articleId, userId)
	if err != nil {
		return fmt.Errorf("can't delete article")
	}
	return nil
}
func (s *OloService) DeleteWidgetForUser(widgetId, userId int64) error {
	err := s.repo.DeleteWidgetForUser(widgetId, userId)
	if err != nil {
		return fmt.Errorf("can't add widget for user")
	}
	return nil
}
