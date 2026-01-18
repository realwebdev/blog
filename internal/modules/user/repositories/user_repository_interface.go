package user

import (
	"github.com/realwebdev/blog/internal/modules/user/models"
)

type UserRepositoryInterface interface {
	RegisterUser(u *models.User) (models.User, error)
	FindByID(id string) (models.User, error)
	FindByEmail(email string) (models.User, error)
}
