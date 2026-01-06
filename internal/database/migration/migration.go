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

	migrationDir := "internal/database/migration"

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

/*

-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS users;


--------------------------------------------articles---------------------


-- +goose Up
CREATE TABLE articles (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS users;




*/
