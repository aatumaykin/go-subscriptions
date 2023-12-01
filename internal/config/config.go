package config

type Config struct {
	Storage string `json:"storage"`
}

func NewConfig() *Config {
	return &Config{
		Storage: "memory",
	}
}
