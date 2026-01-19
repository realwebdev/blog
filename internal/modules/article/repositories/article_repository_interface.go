package repositories

import (
	"github.com/realwebdev/blog/internal/modules/article/models"
)

type ArticleRepositoryInterface interface {
	List(limit int64) []models.Article
	Find(id int64) models.Article // single article return
	Create(article *models.Article) (*models.Article, error)
}
