package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	userRepo "github.com/realwebdev/blog/internal/modules/user/repositories"
	"github.com/realwebdev/blog/internal/modules/user/responses"
	"github.com/realwebdev/blog/pkg/database"
	"github.com/realwebdev/blog/pkg/sessions"
)

func Auth(c *gin.Context) responses.User {
	var response responses.User
	authID := sessions.Get(c, "auth")
	userID, _ := strconv.Atoi(authID)

	var userRepo = userRepo.New(database.Connection())
	user, _ := userRepo.FindByID(userID)
	if user.ID == 0 {
		return response
	}

	return responses.ToUser(user)
}
