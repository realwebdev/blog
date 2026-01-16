package routes

import (
	"github.com/gin-gonic/gin"
	userCtrl "github.com/realwebdev/blog/internal/modules/user/controllers"
	userRepo "github.com/realwebdev/blog/internal/modules/user/repositories"
)

func Routes(router *gin.Engine) {
	userRepostiory := userRepo.NewService()
	userController := userCtrl.New(baseController)
	router.GET("/register", userController.Register)
	router.POST("/register", userController.HandleRegister)
}
