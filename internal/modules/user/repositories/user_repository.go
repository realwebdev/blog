package user

import (
	"errors"
	"net/mail"

	"github.com/realwebdev/blog/internal/modules/user/models"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo UserRepository
}

func NewService(repo UserRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) RegisterUser(name, email, password string) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return errors.New("invalid email format")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}
	return s.repo.Create(user)
}
