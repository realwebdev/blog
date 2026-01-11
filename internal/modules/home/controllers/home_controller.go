package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/realwebdev/blog/internal/modules/article/services"
	"github.com/realwebdev/blog/pkg/html"
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
	html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
		"title":    "Home page",
		"featured": controller.articleService.GetFeaturedArticles(), // TODO: feature keyword is not showing JSON object
		"stories":  controller.articleService.GetStoriesArticles(),
	})

}

