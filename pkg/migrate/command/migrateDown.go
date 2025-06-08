package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"rwa/config"
	"rwa/internal/database"
	"rwa/pkg/migrate"
)

var MigrateDownCmd = &cobra.Command{
	Use:   "migrate:down",
	Short: "Migrate up",
	Long:  "Execute migrate up",
	Run:   down,
}

func down(cmd *cobra.Command, args []string) {
	db, err := database.InitSqlDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	migrator := migrate.NewMigrator(db.GetDB())
	err = migrator.Down(config.MigrationsPath)
	if err != nil {
		fmt.Println(err)
		return
	}
}
