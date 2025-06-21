package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"rwa/pkg/migrate/command"
)

var rootCmd = &cobra.Command{}

func Execute() error {
	AddCommands()
	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return err
		}
		os.Exit(1)
	}

	return nil
}

func AddCommands() {
	rootCmd.AddCommand(command.MigrateUpCmd)
	rootCmd.AddCommand(command.MigrateDownCmd)
	rootCmd.AddCommand(command.CreateCmd)
}
