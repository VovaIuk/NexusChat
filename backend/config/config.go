package config

import (
	"backend/internal/adapter/postgres"
	"backend/pkg/httpserver"
	jwttoken "backend/pkg/jwt_token"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type App struct {
	Name    string `envconfig:"APP_NAME"    required:"true"`
	Version string `envconfig:"APP_VERSION" required:"true"`
}

type Config struct {
	App      App
	HTTP     httpserver.Config
	JWT      jwttoken.Config
	Postgres postgres.Config
}

func InitConfig() (Config, error) {
	// Загружаем .env, если он есть (для локальной разработки)
	_ = godotenv.Load()

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, fmt.Errorf("failed to load config: %v", err)
	}

	return cfg, nil
}
