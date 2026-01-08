package models

import (
	"time"

	"github.com/realwebdev/blog/internal/modules/user/models"
)

type Article struct {
	ID        int64       `json:"id"`
	Title     string      `json:"title"`
	Content   string      `json:"content"`
	UserID    int64       `json:"user_id"`
	User      models.User `json:"user"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
