package seeder

import (
	"log"
	"os"

	"github.com/realwebdev/blog/internal/modules/user"
	"github.com/realwebdev/blog/pkg/database"
)

func Seed() {
	log.Println("Seeding database...")

	repo := &user.PostgresUserRepository{
		DB: database.Connection(),
	}
	service := user.NewService(repo)

	// Create Admin User
	password := os.Getenv("ADMIN_PASSWORD")
	if password == "" {
		password = "password123"
		log.Println("WARNING: Using default insecure password 'password123' for admin. Set ADMIN_PASSWORD env var to override.")
	}

	err := service.RegisterUser("Admin User", "admin@example.com", password)
	if err != nil {
		log.Printf("Error seeding user (might already exist): %v", err)
	} else {
		log.Println("User created successfully")
	}

	log.Println("Seeding complete.")
}
