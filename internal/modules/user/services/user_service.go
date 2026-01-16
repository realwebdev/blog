package services

import (
	"errors"
	"net/mail"

	"github.com/realwebdev/blog/internal/modules/user/models"
	user "github.com/realwebdev/blog/internal/modules/user/repositories"
	"github.com/realwebdev/blog/internal/modules/user/requests/auth"
	"github.com/realwebdev/blog/internal/modules/user/responses"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository user.UserRepositoryInterface
}

func NewUserService(us UserService) *UserService {
	return &UserService{
		userRepository: us.userRepository,
	}

}

func (s *UserService) RegisterUser(request auth.RegisterRequest) (responses.User, error) {
	var response responses.User
	var user models.User

	if _, err := mail.ParseAddress(user.Email); err != nil {
		return responses.User{}, errors.New("invalid email format")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return responses.User{}, err
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Password = string(hashedPassword)

	newUser, err := s.userRepository.RegisterUser(&user)
	if err != nil {
		return response, err
	}

	return responses.ToUser(newUser), nil
}
