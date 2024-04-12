package handler

import (
	"OLO-backend/olo_service/generated"
	"OLO-backend/olo_service/internal/entity"
	"OLO-backend/olo_service/internal/service"
	"OLO-backend/olo_service/internal/utils/jwt"
	"context"
	"errors"
	"fmt"
	jwt_go "github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/metadata"
)

type OloHandler struct {
	service   *service.OloService
	validator *jwt.Validator
	generated.UnimplementedOLOServer
}

func NewOloHandler(service *service.OloService, validator *jwt.Validator) *OloHandler {
	return &OloHandler{
		service:   service,
		validator: validator,
	}
}

func (h *OloHandler) tokenFromContextMetadata(ctx context.Context) (*jwt_go.Token, error) {
	// rip the token from the metadata via the context
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("no metadata found in context")
	}
	tokens := headers.Get("Authorization")
	if len(tokens) < 1 {
		return nil, errors.New("no token found in metadata")
	}
	tokenString := tokens[0]

	token, err := h.validator.GetToken(tokenString)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	return token, nil
}

func (h *OloHandler) newEntityUseToken(token *jwt_go.Token) entity.User {
	id := int64(token.Claims.(jwt_go.MapClaims)["uid"].(float64))
	email := token.Claims.(jwt_go.MapClaims)["email"].(string)
	role := token.Claims.(jwt_go.MapClaims)["role"].(string)

	return entity.User{
		ID:    id,
		Email: email,
		Role:  role,
	}
}

func (h *OloHandler) HelloUser(ctx context.Context, req *generated.HelloUserRequest) (*generated.HelloUserResponse, error) {
	token, err := h.tokenFromContextMetadata(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get token: %w", err)
	}
	user := h.newEntityUseToken(token)
	return &generated.HelloUserResponse{
		Message: fmt.Sprintf(
			"Hello %s (%d)! I am the olo service. You have role %s",
			user.Email, user.ID, user.Role),
	}, nil
}

func (h *OloHandler) mapToProtoModelWidgets(widgets []entity.Widget) []*generated.Widget {
	var result []*generated.Widget
	for _, widget := range widgets {
		result = append(result, &generated.Widget{
			Id:          uint64(widget.ID),
			Description: widget.Description,
		})
	}
	return result
}
func (h *OloHandler) mapToProtoModelArticles(articles []entity.Article) []*generated.Article {
	var result []*generated.Article
	for _, article := range articles {
		result = append(result, &generated.Article{
			Id:     uint64(article.ID),
			Header: article.Header,
		})
	}
	return result
}

func (h *OloHandler) GetAllWidgets(ctx context.Context, req *generated.GetWidgetsRequest) (*generated.GetWidgetsResponse, error) {
	_, err := h.tokenFromContextMetadata(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get token: %w", err)
	}
	widgets, err := h.service.GetAllWidgets()
	if err != nil {
		return nil, err
	}

	return &generated.GetWidgetsResponse{
		Widgets: h.mapToProtoModelWidgets(widgets),
	}, nil
}

func (h *OloHandler) GetUserWidgets(ctx context.Context, req *generated.GetWidgetsRequest) (*generated.GetWidgetsResponse, error) {
	token, err := h.tokenFromContextMetadata(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get token: %w", err)
	}
	user := h.newEntityUseToken(token)
	widgets, err := h.service.GetUserWidgets(user.ID)
	if err != nil {
		return nil, err
	}

	return &generated.GetWidgetsResponse{
		Widgets: h.mapToProtoModelWidgets(widgets),
	}, nil
}

func (h *OloHandler) AddWidgetForUser(ctx context.Context, req *generated.AddWidgetForUserRequest) (*generated.AddWidgetForUserResponse, error) {
	token, err := h.tokenFromContextMetadata(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get token: %w", err)
	}
	user := h.newEntityUseToken(token)
	err = h.service.AddWidgetForUser(req.WidgetId, user.ID)
	if err != nil {
		return nil, fmt.Errorf("could not add widget to user %w", err)
	}
	return &generated.AddWidgetForUserResponse{
		Response: fmt.Sprintf(
			"Successfully add widget (%d) for user (%d)!",
			req.WidgetId, user.ID),
	}, nil
}

func (h *OloHandler) GetUsersArticles(ctx context.Context, req *generated.GetAllArticlesRequest) (*generated.GetAllArticlesResponse, error) {
	token, err := h.tokenFromContextMetadata(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get token: %w", err)
	}
	user := h.newEntityUseToken(token)
	articles, err := h.service.GetUsersArticles(user.ID)
	if err != nil {
		return nil, err
	}

	return &generated.GetAllArticlesResponse{
		Articles: h.mapToProtoModelArticles(articles),
	}, nil
}

func (h *OloHandler) GetAllArticles(ctx context.Context, req *generated.GetAllArticlesRequest) (*generated.GetAllArticlesResponse, error) {
	_, err := h.tokenFromContextMetadata(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get token: %w", err)
	}
	articles, err := h.service.GetAllArticles()
	if err != nil {
		return nil, err
	}

	return &generated.GetAllArticlesResponse{
		Articles: h.mapToProtoModelArticles(articles),
	}, nil
}

func (h *OloHandler) AddArticleForUser(ctx context.Context, req *generated.AddArticleForUserRequest) (*generated.AddArticleForUserResponse, error) {
	token, err := h.tokenFromContextMetadata(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get token: %w", err)
	}
	user := h.newEntityUseToken(token)
	err = h.service.AddArticleForUser(req.ArticleId, user.ID)
	if err != nil {
		return nil, fmt.Errorf("could not add article to user %w", err)
	}
	return &generated.AddArticleForUserResponse{
		Response: fmt.Sprintf(
			"Successfully add article (%d) for user (%d)!",
			req.ArticleId, user.ID),
	}, nil
}
