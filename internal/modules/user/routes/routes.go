package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/realwebdev/blog/internal/modules/user/controllers"
	userRepository "github.com/realwebdev/blog/internal/modules/user/repositories"
	"github.com/realwebdev/blog/internal/modules/user/services"
	"github.com/realwebdev/blog/pkg/database"
)

func Routes(router *gin.Engine) {
	userRepo := userRepository.New(database.Connection())
	userSer := services.NewUserService(userRepo)
	userCtrl := controllers.New(userSer)

	router.GET("/register", userCtrl.Register)
	router.POST("/register", userCtrl.HandleRegister)

	router.GET("/login", userCtrl.Login)
	router.POST("/login", userCtrl.HandleLogin)
}
