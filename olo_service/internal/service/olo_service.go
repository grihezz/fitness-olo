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
	"OLO-backend/pkg/utils/logger/sl"
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

// GetWidgets retrieves widgets associated with a specific user from the repository.
func (s *OloService) GetWidgets(userId int64) ([]entity.Widget, error) {
	const op = "olo.GetWidgets"
	widgets, err := s.repo.GetWidgets(userId)
	if err != nil {
		return nil, sl.Wrap(op, fmt.Errorf("can't get widgets"))
	}
	return widgets, nil
}

// UpdateWidget adds a widget for a specific user.
func (s *OloService) UpdateWidget(data string, widgetId, userId int64) error {
	const op = "olo.UpdateWidget"
	err := s.repo.UpdateWidget(data, widgetId, userId)
	if err != nil {
		return sl.Wrap(op, fmt.Errorf("can't update widget"))
	}
	return nil
}

// AddWidget adds a widget for a specific user.
func (s *OloService) AddWidget(data string, userId int64) (int64, error) {
	const op = "olo.AddWidget"
	widgetId, err := s.repo.AddWidget(data, userId)
	if err != nil {
		return 0, sl.Wrap(op, fmt.Errorf("can't add widget"))
	}
	return widgetId, nil
}

// DeleteWidgetForUser delete a widget for a specific user.
func (s *OloService) DeleteWidgetForUser(widgetId, userId int64) error {
	const op = "olo.DeleteWidgetForUser"
	err := s.repo.DeleteWidget(widgetId, userId)
	if err != nil {
		return sl.Wrap(op, fmt.Errorf("can't delete widget for user"))
	}
	return nil
}

// GetAllArticles retrieves all articles from the repository.
func (s *OloService) GetAllArticles() ([]entity.Article, error) {
	const op = "olo.GetAllArticles"
	articles, err := s.repo.GetAllArticles()
	if err != nil {
		return nil, sl.Wrap(op, fmt.Errorf("can't get all articles"))
	}
	return articles, nil
}

// GetUsersArticles retrieves articles associated with a specific user from the repository.
func (s *OloService) GetUsersArticles(userId int64) ([]entity.Article, error) {
	const op = "olo.GetUsersArticles"
	articles, err := s.repo.GetUsersArticles(userId)
	if err != nil {
		return nil, sl.Wrap(op, fmt.Errorf("can't get articles of user"))
	}
	return articles, nil
}

// AddArticleForUser adds an article for a specific user.
func (s *OloService) AddArticleForUser(articleId, userId int64) error {
	const op = "olo.AddArticleForUser"
	err := s.repo.AddArticleForUser(articleId, userId)
	if err != nil {
		return sl.Wrap(op, fmt.Errorf("can't add article for user"))
	}
	return nil
}

func (s *OloService) DeleteArticleForUser(articleId int64, userId int64) error {
	const op = "olo.DeleteArticleForUser"
	err := s.repo.DeleteArticleForUser(articleId, userId)
	if err != nil {
		return sl.Wrap(op, fmt.Errorf("can't delete article"))
	}
	return nil
}
