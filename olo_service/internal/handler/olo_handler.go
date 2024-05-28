// Package handler provides gRPC handler functions for OLO service endpoints.
//
// This package includes handler functions for handling gRPC requests related to articles and widgets.
package handler

import (
	"OLO-backend/olo_service/generated"
	"OLO-backend/olo_service/internal/entity"
	"OLO-backend/olo_service/internal/mapper"
	"OLO-backend/olo_service/internal/service"
	"OLO-backend/pkg/utils/jwt"
	"context"
	"fmt"
	jwtgo "github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ArticleToArticleResponse converts an Article entity to a generated.Article.
func ArticleToArticleResponse(article entity.Article) *generated.Article {
	return &generated.Article{
		Id:     uint64(article.ID),
		Header: article.Header,
	}
}

// WidgetToWidgetResponse converts a Widget entity to a generated.Widget.
func WidgetToWidgetResponse(widget entity.Widget) *generated.Widget {
	return &generated.Widget{
		Id:   widget.ID,
		Data: widget.Data,
	}
}

// OloHandler represents the gRPC handler for OLO service endpoints.
type OloHandler struct {
	service   *service.OloService
	validator *jwt.Validator

	mapperWidget  mapper.MapFunc[entity.Widget, *generated.Widget]
	mapperArticle mapper.MapFunc[entity.Article, *generated.Article]

	generated.UnimplementedOLOServer
}

func NewOloHandler(service *service.OloService, validator *jwt.Validator) *OloHandler {
	return &OloHandler{
		service:   service,
		validator: validator,

		mapperWidget:  WidgetToWidgetResponse,
		mapperArticle: ArticleToArticleResponse,
	}
}

// getToken retrieves and validates the JWT token from the context.
func (h *OloHandler) getToken(ctx context.Context) (*jwtgo.Token, error) {
	token, err := h.validator.TokenFromContextMetadata(ctx, "Authorization")
	if err != nil {
		// todo wrap error
		return nil, fmt.Errorf("can't get token")
	}
	return token, nil
}

// newEntityUseToken extracts user information from the JWT token.
func (h *OloHandler) newEntityUseToken(token *jwtgo.Token) entity.User {
	id := int64(token.Claims.(jwtgo.MapClaims)["uid"].(float64))
	email := token.Claims.(jwtgo.MapClaims)["email"].(string)
	role := token.Claims.(jwtgo.MapClaims)["role"].(string)

	return entity.User{
		ID:    id,
		Email: email,
		Role:  role,
	}
}

func (h *OloHandler) HelloUser(ctx context.Context, _ *generated.HelloUserRequest) (*generated.HelloUserResponse, error) {
	token, err := h.getToken(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	user := h.newEntityUseToken(token)
	return &generated.HelloUserResponse{
		Message: fmt.Sprintf(
			"Hello %s (%d)! I am the olo service. You have role %s",
			user.Email, user.ID, user.Role),
	}, nil
}

func (h *OloHandler) GetWidgets(ctx context.Context, _ *generated.GetWidgetsRequest) (*generated.GetWidgetsResponse, error) {
	token, err := h.getToken(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	user := h.newEntityUseToken(token)
	widgets, err := h.service.GetWidgets(user.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &generated.GetWidgetsResponse{
		Widgets: h.mapperWidget.MapEach(widgets),
	}, nil
}

func (h *OloHandler) UpdateWidget(ctx context.Context, req *generated.Widget) (*generated.WidgetResponse, error) {
	token, err := h.getToken(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	user := h.newEntityUseToken(token)

	err = h.service.UpdateWidget(req.Data, req.Id, user.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &generated.WidgetResponse{
		Response: fmt.Sprintf(
			"Successfully update widget (%d) for user (%d)!",
			req.Id, user.ID),
	}, nil
}

func (h *OloHandler) AddWidget(ctx context.Context, req *generated.AddWidgetRequest) (*generated.WidgetResponse, error) {
	token, err := h.getToken(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	user := h.newEntityUseToken(token)

	widgetId, err := h.service.AddWidget(req.Data, user.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &generated.WidgetResponse{
		Response: fmt.Sprintf(
			"Successfully add widget (%d) for user (%d)!",
			widgetId, user.ID),
	}, nil
}

func (h *OloHandler) GetUsersArticles(ctx context.Context, _ *generated.GetAllArticlesRequest) (*generated.GetAllArticlesResponse, error) {
	token, err := h.getToken(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	user := h.newEntityUseToken(token)
	articles, err := h.service.GetUsersArticles(user.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &generated.GetAllArticlesResponse{
		Articles: h.mapperArticle.MapEach(articles),
	}, nil
}

func (h *OloHandler) GetAllArticles(ctx context.Context, _ *generated.GetAllArticlesRequest) (*generated.GetAllArticlesResponse, error) {
	_, err := h.getToken(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	articles, err := h.service.GetAllArticles()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &generated.GetAllArticlesResponse{
		Articles: h.mapperArticle.MapEach(articles),
	}, nil
}

func (h *OloHandler) AddArticleForUser(ctx context.Context, req *generated.ArticleForUserRequest) (*generated.ArticleForUserResponse, error) {
	token, err := h.getToken(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	user := h.newEntityUseToken(token)
	err = h.service.AddArticleForUser(req.ArticleId, user.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &generated.ArticleForUserResponse{
		Response: fmt.Sprintf(
			"Successfully add article (%d) for user (%d)!",
			req.ArticleId, user.ID),
	}, nil
}

func (h *OloHandler) DeleteArticleForUser(ctx context.Context, req *generated.ArticleForUserRequest) (*generated.ArticleForUserResponse, error) {
	token, err := h.getToken(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	user := h.newEntityUseToken(token)
	err = h.service.DeleteArticleForUser(req.ArticleId, user.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &generated.ArticleForUserResponse{
		Response: fmt.Sprintf(
			"Successfully delete article (%d) for user (%d)!",
			req.ArticleId, user.ID),
	}, nil
}

func (h *OloHandler) DeleteWidget(ctx context.Context, req *generated.DeleteWidgetRequest) (*generated.WidgetResponse, error) {
	token, err := h.getToken(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	user := h.newEntityUseToken(token)
	err = h.service.DeleteWidgetForUser(req.WidgetId, user.ID)
	if err != nil {
		return nil, err
	}
	return &generated.WidgetResponse{
		Response: fmt.Sprintf(
			"Successfully delete widget (%d) for user (%d)!",
			req.WidgetId, user.ID),
	}, nil
}
