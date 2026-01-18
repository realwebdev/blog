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

func New(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
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

func (r *UserRepository) FindByEmail(email string) (models.User, error) {
	var user models.User
	query := `SELECT id, name, email, password FROM users WHERE email = $1`
	if err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) FindByID(id string) (models.User, error) {
	var user models.User
	query := `SELECT id, name, email FROM users WHERE email = $1`
	if err := r.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email); err != nil {
		return models.User{}, err
	}

	return user, nil
}
