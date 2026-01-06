package repositories

import (
	"database/sql"

	"github.com/realwebdev/blog/internal/modules/article/models"
	"github.com/realwebdev/blog/pkg/database"
)

type ArticleRepository struct {
	DB *sql.DB
}

func New() *ArticleRepository {
	return &ArticleRepository{
		DB: database.Connection(),
	}
}

func (articleRepository *ArticleRepository) List(limit int) []models.Article {
	var article []models.Article

List article

	return article
}

func (articleRepository *ArticleRepository) Find(id int) models.Article {
	var article []models.Article

//Find

	return article
}

func (articleRepository *ArticleRepository) Create(article ArticleModel.Article) ArticleModel.Article {
	var newArticle ArticleModel.Article

	articleRepository // create

	return newArticle
}
