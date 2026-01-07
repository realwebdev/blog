package services

import (
	"github.com/realwebdev/blog/internal/modules/article/models"
)

type ArticleRepostioryInterface interface {
	List(limit int) []models.Article
	Find(id int) models.Article
	Create(article models.Article) models.Article
}
