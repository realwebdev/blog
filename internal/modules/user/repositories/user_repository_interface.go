package user

import (
	"database/sql"
	"time"

	"github.com/realwebdev/blog/internal/modules/user/models"
)

type UserRepository interface {
	Create(u *models.User) error
}

type PostgresUserRepository struct {
	DB *sql.DB
}

func (r *PostgresUserRepository) Create(u *models.User) error {
	query := `INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.DB.Exec(query, u.Name, u.Email, u.Password, time.Now(), time.Now())
	return err
}

