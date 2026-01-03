package repository

import "github.com/realwebdev/blog/internal/models"

type DatabaseRepo interface {
	Connection() interface{}
	CreateUser(user *models.User) (*models.User, error)
	CreateArticle(article *models.Article) (*models.Article, error)
}
