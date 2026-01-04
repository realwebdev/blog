package cmd

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/realwebdev/blog/pkg/bootstrap"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate [command] [args]",
	Short: "Run database migrations",
	Long:  `Run database migrations using goose. Commands: up, down, status, create, redo, reset, version.`,
	Run: func(cmd *cobra.Command, args []string) {
		migrate(args)
	},
}

func migrate(args []string) {
	bootstrap.Migrate(args)

}
