package services

import (
	"github.com/realwebdev/blog/internal/modules/article/repositories"
	"github.com/realwebdev/blog/internal/modules/article/responses"
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

func (as *ArticleService) GetFeaturedArticles() responses.Articles {
	articles := as.articleRepository.List(4)
	return responses.ToArticles(articles)
}
func (as *ArticleService) GetStoriesArticles() responses.Articles {
	articles := as.articleRepository.List(4)
	return responses.ToArticles(articles)
}
