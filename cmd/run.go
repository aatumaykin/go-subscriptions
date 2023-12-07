package cmd

import (
	"git.home/alex/go-subscriptions/internal/api"
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

		httpServer, err := api.NewHTTPServer(
			api.WithTimeout(application.Config.Timeout),
			api.WithListenAddr(application.Config.ListenAddr),
			api.WithContext(application.Context),
			api.WithDefaultRouter(),
			api.WithHealthHandler(),
			api.WithCategoryCreateHandler(application.ServiceFactory.CategoryService),
			api.WithCategoryGetHandler(application.ServiceFactory.CategoryService),
			api.WithCategoryCollectionGetHandler(application.ServiceFactory.CategoryService),
			api.WithCategoryUpdateHandler(application.ServiceFactory.CategoryService),
			api.WithCategoryDeleteHandler(application.ServiceFactory.CategoryService),
		)
		if err != nil {
			return err
		}

		httpServer.ListenAndServe()

		return nil
	},
}
