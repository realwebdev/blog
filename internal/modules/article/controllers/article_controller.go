package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/realwebdev/blog/internal/modules/article/services"
)

type Controller struct {
	articleService services.ArticleServiceInterface
}

func New(aSi services.ArticleServiceInterface) *Controller {
	return &Controller{
		articleService: aSi,
	}
}

func (controller *Controller) Show(c *gin.Context) {
	// TODO:
	// 1. Get the article
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error converting the id"})
		return
	}

	article, err := controller.articleService.Find(int64(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ID": article})
}

func (controller *Controller) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "article create ...",
	})
}
