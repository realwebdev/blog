package cmd

import (
	"log"

	"github.com/realwebdev/blog/internal/models"
	"github.com/realwebdev/blog/internal/repository/dbrepo"
	"github.com/realwebdev/blog/pkg/database"
	"github.com/spf13/cobra"
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed the database with initial data",
	Run: func(cmd *cobra.Command, args []string) {
		database.Connect()
		defer database.DB.Close()

		seed()
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}

func seed() {
	log.Println("Seeding database...")

	// Initialize the repository
	repo := dbrepo.NewPostgresRepo(database.DB)

	// Create Admin User
	user := &models.User{
		Name:  "Admin User",
		Email: "admin@example.com",
	}

	createdUser, err := repo.CreateUser(user)
	if err != nil {
		log.Printf("Error seeding user (might already exist): %v", err)
	} else {
		log.Printf("User created with ID: %d", createdUser.ID)
	}

	log.Println("Seeding complete.")
}
