package cmd

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate [command] [args]",
	Short: "Run database migrations",
	Long:  `Run database migrations using goose. Commands: up, down, status, create, redo, reset, version.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}

		command := args[0]
		dsn := os.Getenv("DB_DSN")
		if dsn == "" {
			dsn = "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable"
		}
		/* 
		dsn := "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable"
		db, err := sql.Open("pgx", dsn)
		*/

		db, err := sql.Open("pgx", dsn)
		if err != nil {
			log.Fatalf("Error opening database: %v", err)
		}
		defer db.Close()

		if err := goose.SetDialect("postgres"); err != nil {
			log.Fatal(err)
		}

		migrationDir := "database/migrations"

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

		if err := goose.Run(command, db, migrationDir, migrationArgs...); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
