package responses

import (
	"fmt"

	"github.com/realwebdev/blog/internal/modules/article/models"
	"github.com/realwebdev/blog/internal/modules/user/responses"
)

// This custom layer will convert the database response to required frontend response. i.e format date, add image and etc.
type Article struct {
	ID        int64
	Image     string
	Title     string
	Content   string
	CreatedAt string
	User      responses.User
}

type Articles struct {
	Data []Article
}

func ToArticle(article models.Article) Article {
	return Article{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Image:     "/assets/img/demopic/10.jpg",
		CreatedAt: fmt.Sprintf("%d/%02d/%02d", article.CreatedAt.Year(), article.CreatedAt.Month(), article.CreatedAt.Day()),
		User:      responses.ToUser(article.User),
	}
}

func ToArticles(articles []models.Article) Articles {
	var response Articles

	for _, article := range articles {
		response.Data = append(response.Data, ToArticle(article))
	}

	return response
}
