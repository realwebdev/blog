package services

import (
	"errors"

	articleModel "github.com/realwebdev/blog/internal/modules/article/models"
	"github.com/realwebdev/blog/internal/modules/article/repositories"
	"github.com/realwebdev/blog/internal/modules/article/requests/articles"
	"github.com/realwebdev/blog/internal/modules/article/responses"
	userResponse "github.com/realwebdev/blog/internal/modules/user/responses"
)

var _ ArticleServiceInterface = (*ArticleService)(nil)

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

func (as *ArticleService) Find(id int64) (responses.Article, error) {
	var response responses.Article

	article := as.articleRepository.Find(id)

	if article.ID == 0 {
		return response, errors.New("article not found")
	}

	return responses.ToArticle(article), nil
}

func (as *ArticleService) Store(request articles.StoreRequest, user userResponse.User) (responses.Article, error) {
	var article articleModel.Article

	article.Title = request.Title
	article.Content = request.Content
	article.UserID = user.ID

	newArticle, err := as.articleRepository.Create(&article)
	if err != nil {
		return responses.ToArticle(article), err
	}

	return responses.ToArticle(*newArticle), nil
}
