package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/realwebdev/blog/internal/modules/article/requests/articles"
	"github.com/realwebdev/blog/internal/modules/article/services"
	"github.com/realwebdev/blog/internal/modules/user/helpers"
	"github.com/realwebdev/blog/pkg/converters"
	"github.com/realwebdev/blog/pkg/errors"
	"github.com/realwebdev/blog/pkg/html"
	"github.com/realwebdev/blog/pkg/old"
	"github.com/realwebdev/blog/pkg/sessions"
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

	// If the article is not found, show error page
	if err != nil {
		html.Render(c, http.StatusNotFound, "templates/errors/html/404", gin.H{"title": "Page not found", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ID": article})
}

func (controller *Controller) Create(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/article/html/create", gin.H{
		"title": "Create article",
	})
}

func (controller *Controller) Store(c *gin.Context) {
	var request articles.StoreRequest

	if err := c.ShouldBind(&request); err != nil {
		validationErrors := errors.FromValidation(err)

		sessions.Set(c, "errors", converters.MapToString(validationErrors))

		sessions.Set(c, "old", converters.UrlValuesToString(old.FromContext(c)))

		c.Redirect(http.StatusFound, "/articles/create")
		return
	}

	user := helpers.Auth(c)

	// create the article
	article, err := controller.articleService.Store(request, user)
	if err != nil {
		sessions.Set(c, "errors", converters.MapToString(map[string]string{
			"article": err.Error(),
		}))
		sessions.Set(c, "old", converters.UrlValuesToString(old.FromContext(c)))

		c.Redirect(http.StatusFound, "/articles/create")
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/articles/%d", article.ID))
}

func (controller *Controller) HandleLogout(c *gin.Context) {
	sessions.Remove(c, "auth")
	c.Redirect(http.StatusFound, "/")
}
