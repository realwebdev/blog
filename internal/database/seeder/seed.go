package seeder

import (
	"fmt"
	"log"
	"os"

	articleModels "github.com/realwebdev/blog/internal/modules/article/models"
	"github.com/realwebdev/blog/internal/modules/article/repositories"
	userModels "github.com/realwebdev/blog/internal/modules/user/models"
	user "github.com/realwebdev/blog/internal/modules/user/repositories"
	"golang.org/x/crypto/bcrypt"

	"github.com/realwebdev/blog/pkg/database"
)

func Seed() {
	log.Println("Seeding database...")

	db := database.Connection()
	userRepo := user.New(db)
	articleRepo := repositories.New(db)

	// Define users to seed
	adminPassword := os.Getenv("ADMIN_PASSWORD")
	if adminPassword == "" {
		adminPassword = "password123"
		log.Println("WARNING: Using default insecure password 'password123' for admin. Set ADMIN_PASSWORD env var to override.")
	}

	usersToSeed := []struct {
		Name     string
		Email    string
		Password string
	}{
		{"Admin User", "admin@example.com", adminPassword},
		{"John Doe", "john@example.com", "secret"},
		{"Jane Smith", "jane@example.com", "secret"},
		{"Bob Wilson", "bob@example.com", "secret"},
		{"Alice Brown", "alice@example.com", "secret"},
	}

	// aRepo := &articleRepo.ArticleRepository{DB: db}

	for _, u := range usersToSeed {
		// Register User
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		newUser := &userModels.User{
			Name:     u.Name,
			Email:    u.Email,
			Password: string(hashedPassword),
		}
		user, err := userRepo.RegisterUser(newUser)
		if err != nil {
			log.Printf("Error seeding user %s (might already exist): %v", u.Name, err)
		} else {
			log.Printf("User %s created successfully", u.Name)
		}

		// Seed 6 Articles per user
		for i := 1; i <= 6; i++ {
			article := articleModels.Article{
				Title:   fmt.Sprintf("%s - Article %d", u.Name, i),
				Content: fmt.Sprintf("This is the content for article %d written by %s.\n\nLorem ipsum dolor sit amet.", i, u.Name),
				UserID:  user.ID,
			}

			if err := articleRepo.Create(&article); err != nil {
				log.Printf("Error seeding article '%s': %v", article.Title, err)
			} else {
				log.Printf("Article '%s' seeded successfully", article.Title)
			}
		}
	}

	log.Println("Seeding complete.")
}
