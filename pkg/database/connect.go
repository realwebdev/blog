package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func Connect() {
	var err error
	dsn := os.Getenv("DB_DSN")

	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/blog_db?sslmode=disable"
	}

	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Database connected successfully")
}
