package cmd

import (
	"os"

	application "git.home/alex/go-subscriptions/internal/app"
	"github.com/spf13/cobra"
)

var app *application.App

var rootCmd = &cobra.Command{
	Use: "subscriptions",
}

func Execute() {
	initCommands()

	app = application.NewApp()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initCommands() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(versionCmd)
}
