package repositories

import (
	"database/sql"
	"log"

	"github.com/realwebdev/blog/internal/modules/article/models"
	userModels "github.com/realwebdev/blog/internal/modules/user/models"
)

// This line forces the compiler to check if ArticleRepository implements ArticleRepositoryInterface
var _ ArticleRepositoryInterface = (*ArticleRepository)(nil)

type ArticleRepository struct {
	DB *sql.DB
}

func New(db *sql.DB) *ArticleRepository {
	return &ArticleRepository{
		DB: db,
	}
}

// TODO: return errors.
func (r *ArticleRepository) List(limit int64) []models.Article {
	var articles []models.Article
	query := `
		SELECT a.id, a.title, a.content, a.user_id, u.id, u.name, u.email
		FROM articles a
		LEFT JOIN users u ON a.user_id = u.id
		LIMIT $1
	`
	rows, err := r.DB.Query(query, limit)
	if err != nil {
		log.Println("Error fetching articles:", err)
		return articles
	}
	defer rows.Close()

	for rows.Next() {
		var article models.Article
		var user userModels.User
		if err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.UserID, &user.ID, &user.Name, &user.Email); err != nil {
			log.Println("Error scanning article:", err)
			continue
		}
		article.User = user
		articles = append(articles, article)
	}

	return articles
}

func (r *ArticleRepository) Find(id int64) models.Article {
	var article models.Article
	var user userModels.User

	query := `
		SELECT a.id, a.title, a.content, a.user_id, u.id, u.name, u.email
		FROM articles a
		LEFT JOIN users u ON a.user_id = u.id
		WHERE a.id = $1
	`
	err := r.DB.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Content, &article.UserID, &user.ID, &user.Name, &user.Email)
	if err != nil {
		log.Println("Error finding article:", err)
		return models.Article{}
	}

	article.User = user
	return article
}

func (r *ArticleRepository) Create(article *models.Article) (*models.Article, error) {
	query := `
		INSERT INTO articles (title, content, user_id, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
		RETURNING *
	`

	err := r.DB.QueryRow(query, article.Title, article.Content, article.UserID).Scan(
		&article.ID,
		&article.Title,
		&article.Content,
		&article.UserID,
		&article.CreatedAt,
		&article.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return article, nil
}
