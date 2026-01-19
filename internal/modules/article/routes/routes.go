package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/realwebdev/blog/internal/middleware"
	"github.com/realwebdev/blog/internal/modules/article/controllers"
	"github.com/realwebdev/blog/internal/modules/article/repositories"
	"github.com/realwebdev/blog/internal/modules/article/services"
	"github.com/realwebdev/blog/pkg/database"
)

func Routes(router *gin.Engine) {
	articleRepository := repositories.New(database.Connection())
	articleService := services.New(articleRepository)
	articleController := controllers.New(articleService)
	router.GET("/articles/:id", articleController.Show)

	authGroup := router.Group("/articles")
	authGroup.Use(middleware.IsAuth())
	{
		authGroup.GET("/create", articleController.Create)
		authGroup.POST("/store", articleController.Store)
	}

}
