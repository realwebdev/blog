package routes

import (
	"github.com/gin-gonic/gin"
	articleRoutes "github.com/realwebdev/blog/internal/modules/article/routes"
	homeRoutes "github.com/realwebdev/blog/internal/modules/home/routes"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
	articleRoutes.Routes(router)
}
