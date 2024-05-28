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

	log := s.log.With(
		slog.String("op", op),
		slog.Int64("userId", userId))

	widgets, err := s.repo.GetWidgets(userId)
	if err != nil {
		log.Error("failed get widgets", sl.Err(err))
		return nil, sl.Wrap(op, fmt.Errorf("can't get widgets"))
	}
	return widgets, nil
}

// UpdateWidget adds a widget for a specific user.
func (s *OloService) UpdateWidget(data string, widgetId, userId int64) error {
	const op = "olo.UpdateWidget"

	log := s.log.With(
		slog.String("op", op),
		slog.Int64("userId", userId),
		slog.Int64("widgetId", widgetId))

	err := s.repo.UpdateWidget(data, widgetId, userId)
	if err != nil {
		log.Error("failed update widget", sl.Err(err))
		return sl.Wrap(op, fmt.Errorf("can't update widget"))
	}
	return nil
}

// AddWidget adds a widget for a specific user.
func (s *OloService) AddWidget(data string, userId int64) (int64, error) {
	const op = "olo.AddWidget"

	log := s.log.With(
		slog.String("op", op),
		slog.Int64("userId", userId))

	widgetId, err := s.repo.AddWidget(data, userId)
	if err != nil {
		log.Error("failed add widget", sl.Err(err))
		return 0, sl.Wrap(op, fmt.Errorf("can't add widget"))
	}
	return widgetId, nil
}

// DeleteWidgetForUser delete a widget for a specific user.
func (s *OloService) DeleteWidgetForUser(widgetId, userId int64) error {
	const op = "olo.DeleteWidgetForUser"

	log := s.log.With(
		slog.String("op", op),
		slog.Int64("userId", userId))

	err := s.repo.DeleteWidget(widgetId, userId)
	if err != nil {
		log.Error("failed delete widget", sl.Err(err))
		return sl.Wrap(op, fmt.Errorf("can't delete widget for user"))
	}
	return nil
}

// GetAllArticles retrieves all articles from the repository.
func (s *OloService) GetAllArticles() ([]entity.Article, error) {
	const op = "olo.GetAllArticles"

	log := s.log.With(
		slog.String("op", op))

	articles, err := s.repo.GetAllArticles()
	if err != nil {
		log.Error("failed get articles", sl.Err(err))
		return nil, sl.Wrap(op, fmt.Errorf("can't get all articles"))
	}
	return articles, nil
}

// GetUsersArticles retrieves articles associated with a specific user from the repository.
func (s *OloService) GetUsersArticles(userId int64) ([]entity.Article, error) {
	const op = "olo.GetUsersArticles"

	log := s.log.With(
		slog.String("op", op),
		slog.Int64("userId", userId))

	articles, err := s.repo.GetUsersArticles(userId)
	if err != nil {
		log.Error("failed get articles of user", sl.Err(err))
		return nil, sl.Wrap(op, fmt.Errorf("can't get articles of user"))
	}
	return articles, nil
}

// AddArticleForUser adds an article for a specific user.
func (s *OloService) AddArticleForUser(articleId, userId int64) error {
	const op = "olo.AddArticleForUser"

	log := s.log.With(
		slog.String("op", op),
		slog.Int64("articleId", articleId),
		slog.Int64("userId", userId))

	err := s.repo.AddArticleForUser(articleId, userId)
	if err != nil {
		log.Error("failed add article of user", sl.Err(err))
		return sl.Wrap(op, fmt.Errorf("can't add article for user"))
	}
	return nil
}

func (s *OloService) DeleteArticleForUser(articleId int64, userId int64) error {
	const op = "olo.DeleteArticleForUser"

	log := s.log.With(
		slog.String("op", op),
		slog.Int64("articleId", articleId),
		slog.Int64("userId", userId))

	err := s.repo.DeleteArticleForUser(articleId, userId)
	if err != nil {
		log.Error("failed delete article", sl.Err(err))
		return sl.Wrap(op, fmt.Errorf("can't delete article"))
	}
	return nil
}
