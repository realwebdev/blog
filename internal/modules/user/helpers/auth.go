package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/realwebdev/blog/internal/modules/user/responses"
	"github.com/realwebdev/blog/pkg/sessions"
)

func Auth(c *gin.Context) responses.User {
	authID := sessions.Get(c, "auth")
	userID, _ := strconv.Atoi(authID)
}
