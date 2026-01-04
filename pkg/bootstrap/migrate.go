package bootstrap

import (
	"github.com/realwebdev/blog/internal/database/migration"
	"github.com/realwebdev/blog/pkg/config"
	"github.com/realwebdev/blog/pkg/database"
)

func Migrate(args []string) {
	config.Set()

	database.Connect()

	migration.Migrate(args)
}
