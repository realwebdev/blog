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
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "error converting the id"})
		return
	}

	// 2. Find the article from the database
	article, err := controller.articleService.Find(int64(id))

	// 3. If the article not found, show error page

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "error converting the id"})
		return
	}
	// 4. If article found, render the article page

	// html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
	// 	"title":    "Home page",
	// 	"featured": controller.articleService.GetFeaturedArticles(),
	// 	"stories":  controller.articleService.GetStoriesArticles(),
	// })

	c.JSON(http.StatusOK, gin.H{"ID": id})
}
