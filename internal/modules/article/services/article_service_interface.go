package services

import (
	"github.com/realwebdev/blog/internal/modules/article/models"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() []models.Article
	GetStoriesArticles() []models.Article
}
