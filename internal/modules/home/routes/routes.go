package routes

import (
	"github.com/gin-gonic/gin"
	articleRepositories "github.com/realwebdev/blog/internal/modules/article/repositories"
	articleServices "github.com/realwebdev/blog/internal/modules/article/services"
	homeCtrl "github.com/realwebdev/blog/internal/modules/home/controllers"
)

func Routes(router *gin.Engine) {
	articleRepository := articleRepositories.New()
	articleService := articleServices.New(articleRepository)
	homeController := homeCtrl.New(articleService)
	router.GET("/", homeController.Index)
}
