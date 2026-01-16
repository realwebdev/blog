package services

import (
	"github.com/realwebdev/blog/internal/modules/user/requests/auth"
	"github.com/realwebdev/blog/internal/modules/user/responses"
)

type UserServiceInterface interface {
	RegisterUser(request auth.RegisterRequest) (responses.User, error)
}
