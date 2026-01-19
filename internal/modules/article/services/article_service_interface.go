package services

import (
	"github.com/realwebdev/blog/internal/modules/article/requests/articles"
	"github.com/realwebdev/blog/internal/modules/article/responses"
	userResponse "github.com/realwebdev/blog/internal/modules/user/responses"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() responses.Articles
	GetStoriesArticles() responses.Articles
	Find(id int64) (responses.Article, error)
	Store(request articles.StoreRequest, user userResponse.User) (responses.Article, error)
}
