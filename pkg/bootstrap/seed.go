package bootstrap

import (
	"github.com/realwebdev/blog/internal/database/seeder"
	"github.com/realwebdev/blog/pkg/config"
	"github.com/realwebdev/blog/pkg/database"
)

func Seed() {
	config.Set()

	database.Connect()

	seeder.Seed()
}
