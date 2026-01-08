package responses

import (
	"fmt"

	"github.com/realwebdev/blog/internal/modules/user/models"
)

type User struct {
	ID    int64
	Image string
	Name  string
	Email string
}

type Users struct {
	Data []User
}

func ToUser(user models.User) User {
	return User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Image: fmt.Sprintf("https://ui-avatars.com/api/?name=%s", user.Name),
	}
}
