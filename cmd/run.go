package cmd

import (
	"git.home/alex/go-subscriptions/internal/app"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use: "run",
	RunE: func(cmd *cobra.Command, args []string) error {
		application, err := app.NewApp(cfgFile)
		if err != nil {
			return err
		}

		api := application.NewAPI()
		api.InitRoutes()

		api.ListenAndServe()

		return nil
	},
}
