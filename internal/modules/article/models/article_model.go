package models

import "github.com/realwebdev/blog/internal/modules/user/models"

type Article struct {
	Title   string
	Content string
	ID      int64 `json:"id"` // user id ** not article id
	User    models.User
}
