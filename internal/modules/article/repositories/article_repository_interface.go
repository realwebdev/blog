package repositories

import (
	"github.com/realwebdev/blog/internal/modules/article/models"
)

type ArticleRepositoryInterface interface {
	List(limit int) []models.Article
	Find(id int) []models.Article
	Create(article *models.Article) error
}
