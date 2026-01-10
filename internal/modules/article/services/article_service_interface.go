package services

import "github.com/realwebdev/blog/internal/modules/article/responses"

type ArticleServiceInterface interface {
	GetFeaturedArticles() responses.Articles
	GetStoriesArticles() responses.Articles
	Find(id int64)(responses.Article, error) 
}
