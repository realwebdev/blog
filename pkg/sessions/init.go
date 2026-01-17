package sessions

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Start(router *gin.Engine) {
	store := cookie.NewStore([]byte("secret")) // TODO: remove hardcoded values
	router.Use(sessions.Sessions("mysession", store))
}
