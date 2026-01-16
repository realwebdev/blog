package user

import (
	"database/sql"
	"time"

	"github.com/realwebdev/blog/internal/modules/user/models"
)

var _ UserRepositoryInterface = (*UserRepository)(nil)

type UserRepository struct {
	DB *sql.DB
}

func NewService(repo UserRepository) *UserRepository {
	return &UserRepository{
		DB: repo.DB,
	}
}

func (r *UserRepository) RegisterUser(u *models.User) (models.User, error) {
	query := `INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.DB.Exec(query, u.Name, u.Email, u.Password, time.Now(), time.Now())
	if err != nil {
		return models.User{}, err
	}
	return *u, nil
}
