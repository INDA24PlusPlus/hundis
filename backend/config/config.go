package config

import "github.com/caarlos0/env/v11"

type Configuration struct {
	BaseURL string `env:"BASE_URL"`

	DBHost     string `env:"DB_HOST"`
	DBPort     int    `env:"DB_PORT"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`

	GitHubClientID     string `env:"GITHUB_CLIENT_ID"`
	GitHubClientSecret string `env:"GITHUB_CLIENT_SECRET"`

	JWTSecret string `env:"JWT_SECRET"`
}

var config Configuration

func Init() {
	err := env.Parse(&config)
	if err != nil {
		panic(err)
	}
}

func Config() Configuration {
	return config
}
