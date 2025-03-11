package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"rwa/config"
	"rwa/internal/database"
	"rwa/pkg/migrate"
)

var CreateCmd = &cobra.Command{
	Use:   "migrate:create [fileName]",
	Short: "Create migration",
	Args:  cobra.MinimumNArgs(1),
	Run:   create,
}

func create(cmd *cobra.Command, args []string) {
	db, err := database.InitSqlDB()
	if err != nil {
		return
	}
	migrator := migrate.NewMigrator(db.GetDB())
	err = migrator.Create(config.MigrationsPath, args[0])
	if err != nil {
		fmt.Println(err)
		return
	}
}
