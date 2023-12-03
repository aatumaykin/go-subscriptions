package app

import (
	"git.home/alex/go-subscriptions/internal/api"
	"git.home/alex/go-subscriptions/internal/config"
)

type App struct {
	config *config.Config
}

func NewApp(configFile string) (*App, error) {
	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		return nil, err
	}

	return &App{
		config: cfg,
	}, nil
}

func (a *App) NewAPI() *api.API {
	return api.NewAPI(
		a.config.ListenAddr,
		a.config.Timeout,
	)
}
