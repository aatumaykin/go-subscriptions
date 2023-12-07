package cmd

import (
	"git.home/alex/go-subscriptions/internal/api"
	"git.home/alex/go-subscriptions/internal/api/handler/subscription_handler"
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
			api.WithCategoryHandlers(application.ServiceFactory.CategoryService),
			api.WithCurrencyHandlers(application.ServiceFactory.CurrencyService),
			api.WithCycleHandlers(application.ServiceFactory.CycleService),
			api.WithSubscribeHandlers(&subscription_handler.HandlerOpts{
				SubscriptionService: application.ServiceFactory.SubscriptionService,
				CategoryService:     application.ServiceFactory.CategoryService,
				CycleService:        application.ServiceFactory.CycleService,
				CurrencyService:     application.ServiceFactory.CurrencyService,
			}),
		)
		if err != nil {
			return err
		}

		httpServer.ListenAndServe()

		return nil
	},
}
