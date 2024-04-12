package mapper

import (
	"OLO-backend/olo_service/generated"
	"OLO-backend/olo_service/internal/entity"
)

type Mapper interface {
	WidgetsProtoToEntity([]entity.Widget) []*generated.Widget
	ArticlesProtoToEntity([]entity.Article) []*generated.Article
}

type MapperImpl struct{}

func NewMapperImpl() MapperImpl {
	return MapperImpl{}
}

func (m MapperImpl) WidgetsProtoToEntity(widgets []entity.Widget) []*generated.Widget {
	var result []*generated.Widget
	for _, widget := range widgets {
		result = append(result, &generated.Widget{
			Id:          uint64(widget.ID),
			Description: widget.Description,
		})
	}
	return result
}

func (m MapperImpl) ArticlesProtoToEntity(articles []entity.Article) []*generated.Article {
	var result []*generated.Article
	for _, article := range articles {
		result = append(result, &generated.Article{
			Id:     uint64(article.ID),
			Header: article.Header,
		})
	}
	return result
}
