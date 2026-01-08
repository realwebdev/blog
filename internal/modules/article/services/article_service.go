package services

import (
	"github.com/realwebdev/blog/internal/modules/article/models"
	"github.com/realwebdev/blog/internal/modules/article/repositories"
)

// This service is a middle between controller and the repository. It will call function from repository
// file.
type ArticleService struct { //
	articleRepository repositories.ArticleRepositoryInterface
}

func New(repo repositories.ArticleRepositoryInterface) *ArticleService {
	return &ArticleService{
		articleRepository: repo,
	}
}

func (as *ArticleService) GetFeaturedArticles() []models.Article {

	return as.articleRepository.List(4)
}

func (as *ArticleService) GetStoriesArticles() []models.Article {

	return as.articleRepository.List(4)
}
