package user

import (
	"github.com/realwebdev/blog/internal/modules/user/models"
)

type UserRepositoryInterface interface {
	RegisterUser(u *models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
}
