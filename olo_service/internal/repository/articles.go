package repository

import (
	"OLO-backend/olo_service/internal/entity"
	"OLO-backend/olo_service/internal/repository/provider"
	"fmt"
)

type ArticleRepo struct {
	mysqlProvider *provider.MySQLProvider
}

func (r *ArticleRepo) GetUsersArticles(userId int64) ([]entity.Article, error) {
	return r.getArticles(fmt.Sprintf("SELECT w.id as id, w.header as `header` FROM articles as w, user_has_articles as u WHERE w.id = u.id_articles AND u.id_user = '%d'", userId))
}

func (r *ArticleRepo) AddArticleForUser(articleId, userId int64) error {
	driver, err := r.mysqlProvider.Driver()
	if err != nil {
		return err
	}
	_, err = driver.NamedExec("INSERT INTO `user_has_articles` (`id_articles`, `id_user`) VALUES (:id_articles, :id_user)", map[string]interface{}{
		"id_articles": articleId,
		"id_user":     userId,
	})
	if err != nil {
		return fmt.Errorf("error add article for user: %w", err)
	}
	return err
}

func NewArticleRepo(mysqlProvider *provider.MySQLProvider) *ArticleRepo {
	return &ArticleRepo{mysqlProvider: mysqlProvider}
}

func (r *ArticleRepo) GetAllArticles() ([]entity.Article, error) {
	return r.getArticles(`SELECT * FROM articles`)
}

func (r *ArticleRepo) getArticles(articleQuery string) ([]entity.Article, error) {
	driver, err := r.mysqlProvider.Driver()
	if err != nil {
		return nil, err
	}

	rows, err := driver.Queryx(articleQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []entity.Article
	for rows.Next() {
		var article entity.Article
		if err := rows.StructScan(&article); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}
