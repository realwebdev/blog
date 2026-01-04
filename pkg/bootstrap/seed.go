package bootstrap

import (
	"github.com/realwebdev/blog/internal/database"
	"github.com/realwebdev/blog/internal/database/seeder"
	"github.com/realwebdev/blog/pkg/config"
)

func Seed() {
	config.Set()

	database.Connect()

	seeder.Seed()
}
