package app

import (
	"git.home/alex/go-subscriptions/internal/config"
)

const version = "0.0.1"

type App struct {
	Version string
}

func NewApp() *App {
	_ = config.NewConfig()

	return &App{
		Version: version,
	}
}
