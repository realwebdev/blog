package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/realwebdev/blog/internal/modules/article/services"
)

type Controller struct {
	articleService services.ArticleServiceInterface
}

func New(aSI services.ArticleServiceInterface) *Controller {
	return &Controller{
		articleService: aSI,
	}
}

func (controller *Controller) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"articles": controller.articleService.GetFeaturedArticles(),
		"stories":  controller.articleService.GetStoriesArticles(),
	})
}

// func New(repo repositories.ArticleRepositoryInterface) *ArticleService {
// 	return &ArticleService{
// 		articleRepository: repo,
// 	}
// }
