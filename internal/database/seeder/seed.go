package seeder

import (
	"log"

	"github.com/realwebdev/blog/internal/modules/home/user"
	"github.com/realwebdev/blog/pkg/database"
)

func Seed() {
	log.Println("Seeding database...")

	repo := &user.PostgresUserRepository{
		DB: database.Connection(),
	}
	service := user.NewService(repo)

	// Create Admin User
	err := service.RegisterUser("Admin User", "admin@example.com", "password123")
	if err != nil {
		log.Printf("Error seeding user (might already exist): %v", err)
	} else {
		log.Println("User created successfully")
	}

	log.Println("Seeding complete.")
}
