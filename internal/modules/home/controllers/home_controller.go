package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/realwebdev/blog/internal/modules/article/repositories"
)

type Controller struct {
	articleRepository *repositories.ArticleRepository
}

func New() *Controller {
	return &Controller{
		articleRepository: repositories.New(),
	}
}

func (controller *Controller) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"articles": controller.articleRepository.List(3),
	})
}
