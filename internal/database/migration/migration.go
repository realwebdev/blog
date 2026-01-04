package migration

import (
	"context"
	"log"

	"github.com/pressly/goose/v3"
	"github.com/realwebdev/blog/pkg/database"
)

func Migrate(args []string) {
	if len(args) == 0 {
		log.Fatal("Migration command required. Usage: migrate <command> [args]")
	}

	command := args[0]

	db := database.Connection()

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	migrationDir := "pkg/database"

	if command == "create" {
		if len(args) < 2 {
			log.Fatal("Migration name required. Usage: migrate create <name>")
		}
		if err := goose.Create(db, migrationDir, args[1], "sql"); err != nil {
			log.Fatal(err)
		}
		return
	}

	// Pass remaining args to goose (e.g., for up-to or down-to)
	var migrationArgs []string
	if len(args) > 1 {
		migrationArgs = args[1:]
	}

	if err := goose.RunContext(context.Background(), command, db, migrationDir, migrationArgs...); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

}
