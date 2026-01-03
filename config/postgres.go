package dbrepo

import (
	"database/sql"
	"time"

	"github.com/realwebdev/blog/internal/models"
	"github.com/realwebdev/blog/internal/repository"
)

type PostgresRepo struct {
	DB *sql.DB
}

func NewPostgresRepo(conn *sql.DB) repository.DatabaseRepo {
	return &PostgresRepo{
		DB: conn,
	}
}

func (m *PostgresRepo) Connection() interface{} {
	return m.DB
}

func (m *PostgresRepo) CreateUser(u *models.User) (*models.User, error) {
	query := `
		INSERT INTO users (name, email, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	err := m.DB.QueryRow(query, u.Name, u.Email, time.Now(), time.Now()).Scan(&u.ID)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (m *PostgresRepo) CreateArticle(a *models.Article) (*models.Article, error) {
	query := `
		INSERT INTO articles (title, content, user_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	err := m.DB.QueryRow(
		query,
		a.Title,
		a.Content,
		a.UserID,
		time.Now(),
		time.Now(),
	).Scan(&a.ID)
	if err != nil {
		return nil, err
	}
	return a, nil
}
