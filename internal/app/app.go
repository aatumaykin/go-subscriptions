package app

import (
	"context"
	"errors"

	"git.home/alex/go-subscriptions/internal/config"
	"git.home/alex/go-subscriptions/internal/factory"
)

var (
	errUndefinedStorage = errors.New("undefined storage")
)

type App struct {
	Context        context.Context
	Config         *config.Config
	ServiceFactory *factory.ServiceFactory
}

type Configuration func(a *App) error

func NewApp(configFile string) (*App, error) {
	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		return nil, err
	}

	rf, err := factoryRepository(cfg.Storage)
	if err != nil {
		return nil, err
	}

	sf, err := factory.NewServiceFactory(
		factory.WithRepositoryFactory(rf),
		factory.WithCategoryService(),
		factory.WithCurrencyService(),
		factory.WithCycleService(),
		factory.WithSubscriptionService(),
	)
	if err != nil {
		return nil, err
	}

	a, err := newApp(
		withConfig(cfg),
		withContext(context.Background()),
		withServiceFactory(sf),
	)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func newApp(cfgs ...Configuration) (*App, error) {
	a := &App{}

	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		err := cfg(a)
		if err != nil {
			return nil, err
		}
	}

	return a, nil
}

func withConfig(config *config.Config) Configuration {
	return func(a *App) error {
		a.Config = config
		return nil
	}
}

func withContext(ctx context.Context) Configuration {
	return func(a *App) error {
		a.Context = ctx
		return nil
	}
}

func withServiceFactory(factory *factory.ServiceFactory) Configuration {
	return func(a *App) error {
		a.ServiceFactory = factory
		return nil
	}
}

func factoryRepository(storage string) (*factory.RepositoryFactory, error) {
	if storage == "memory" {
		return factory.NewRepositoryFactory(factory.WithMemoryRepository())
	}

	return nil, errUndefinedStorage
}
