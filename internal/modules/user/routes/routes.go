package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/realwebdev/blog/internal/middleware"
	"github.com/realwebdev/blog/internal/modules/user/controllers"
	userRepository "github.com/realwebdev/blog/internal/modules/user/repositories"
	"github.com/realwebdev/blog/internal/modules/user/services"
	"github.com/realwebdev/blog/pkg/database"
)

func Routes(router *gin.Engine) {
	userRepo := userRepository.New(database.Connection())
	userSer := services.NewUserService(userRepo)
	userCtrl := controllers.New(userSer)

	guestGroup := router.Group("/")
	guestGroup.Use(middleware.IsGuest())
	{
		guestGroup.GET("/register", userCtrl.Register)
		guestGroup.POST("/register", userCtrl.HandleRegister)

		guestGroup.GET("/login", userCtrl.Login)
		guestGroup.POST("/login", userCtrl.HandleLogin)
	}

	authGroup := router.Group("/")
	authGroup.Use(middleware.IsAuth())
	{
		authGroup.POST("/logout", userCtrl.HandleLogout)
	}

}
