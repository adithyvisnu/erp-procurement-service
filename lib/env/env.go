package env

import (
	"context"

	env "github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type ENV struct {
	// SERVER SET UP
	AppPath string `env:"APP_PATH"`
	AppPort string `env:"APP_PORT"`
	// POSTGRE SET UP
	PostgreHost         string `env:"POSTGRES_HOST"`
	PostgrePort         string `env:"POSTGRES_PORT"`
	PostgreUser         string `env:"POSTGRES_USERNAME"`
	PostgrePassword     string `env:"POSTGRES_PASSWORD"`
	PostgreDatabaseName string `env:"POSTGRES_DATABASE_NAME"`
	PostgreSSLMode      string `env:"POSTGRES_SSLMODE"`
	PostgreTimeZone     string `env:"TZ"`
	PostgreSchema       string `env:"POSTGRES_SCHEMA"`
}

func LoadEnv(_ context.Context) ENV {
	_ = godotenv.Load(".env")

	var c ENV
	_ = env.Parse(&c)
	return c
}
