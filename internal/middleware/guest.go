package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	userRepo "github.com/realwebdev/blog/internal/modules/user/repositories"
	"github.com/realwebdev/blog/pkg/database"
	"github.com/realwebdev/blog/pkg/sessions"
)

func IsGuest() gin.HandlerFunc {
	var userRepo = userRepo.New(database.Connection())

	return func(c *gin.Context) {
		authID := sessions.Get(c, "auth")
		userID, _ := strconv.Atoi(authID)

		user, _ := userRepo.FindByID(userID)

		if user.ID != 0 {
			c.Redirect(http.StatusFound, "/") // homepage
			return
		}

		c.Next()
	}
}
