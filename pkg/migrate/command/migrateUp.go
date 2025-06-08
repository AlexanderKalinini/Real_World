package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"rwa/config"
	"rwa/internal/database"
	"rwa/pkg/migrate"
)

var MigrateUpCmd = &cobra.Command{
	Use:   "migrate:up",
	Short: "Migrate up",
	Long:  "Execute migrate up",
	Run:   up,
}

func up(cmd *cobra.Command, args []string) {
	db, err := database.InitSqlDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	migrator := migrate.NewMigrator(db.GetDB())
	err = migrator.Up(config.MigrationsPath)
	if err != nil {
		fmt.Println(err)
		return
	}
}
