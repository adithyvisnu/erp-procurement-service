package env

import (
	"context"
	"fmt"

	env "github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

// type PostgresENV struct {
// 	Host         string `env:"POSTGRES_HOST"`
// 	Port         string `env:"POSTGRES_PORT"`
// 	User         string `env:"POSTGRES_USER"`
// 	Password     string `env:"POSTGRES_PASSWORD"`
// 	DatabaseName string `env:"POSTGRES_DATABASE_NAME"`
// 	SSLMode      string `env:"POSTGRES_SSLMODE"`
// 	TimeZone     string `env:"TZ"`
// }

// type AppENV struct {
// 	Path string `env:"APP_PATH"`
// 	Port string `env:"APP_PORT"`
// }

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
}

func LoadEnv(_ context.Context) ENV {
	_ = godotenv.Load(".env")

	var c ENV
	_ = env.Parse(&c)
	fmt.Printf("%+v", c)
	return c
}
