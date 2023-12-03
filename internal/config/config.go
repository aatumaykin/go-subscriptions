package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Storage    string        `yaml:"storage" env-default:"memory"`
	ListenAddr string        `yaml:"listen_addr" required:"true"`
	Timeout    time.Duration `yaml:"timeout" env-default:"15"`
}

func LoadConfig(configFile string) (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(configFile, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
