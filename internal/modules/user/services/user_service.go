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

func NewUserService(uRI user.UserRepositoryInterface) *UserService {
	return &UserService{
		userRepository: uRI,
	}

}

func (s *UserService) RegisterUser(request auth.RegisterRequest) (responses.User, error) {
	var response responses.User
	var user models.User

	if _, err := mail.ParseAddress(request.Email); err != nil {
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

func (s *UserService) CheckUserExist(email string) (bool, error) {
	_, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *UserService) HandleUserLogin(request auth.LoginRequest) (responses.User, error) {
	var response responses.User

	exist, err := s.userRepository.FindByEmail(request.Email)
	if err != nil {
		return response, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(exist.Password), []byte(request.Password))
	if err != nil {
		return response, errors.New("invalid credentials")
	}

	return responses.ToUser(exist), nil
}
