package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/realwebdev/blog/internal/modules/article/controllers"
	"github.com/realwebdev/blog/internal/modules/article/repositories"
	"github.com/realwebdev/blog/internal/modules/article/services"
)

func Routes(router *gin.Engine) {
	articleRepository := repositories.New()
	articleService := services.New(articleRepository)
	articleController := controllers.New(articleService)
	router.GET("/articles/:id", articleController.Show)
}
