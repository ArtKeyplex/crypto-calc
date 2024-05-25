package configs

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Config struct {
	Port         string
	LogLevel     string `envconfig:"LOG_LEVEL" default:"debug"`
	DSN          string
	FastForexAPI string `envconfig:"FAST_FOREX_API"`
	FastForexURL string `envconfig:"FAST_FOREX_URL"`
}

func LoadConfig() (Config, error) {
	cnf := Config{}

	if err := godotenv.Load(".env"); err != nil && !errors.Is(err, os.ErrNotExist) {
		return cnf, errors.Wrap(err, "read .env file")
	}

	if err := envconfig.Process("", &cnf); err != nil {
		return cnf, errors.Wrap(err, "read environment")
	}

	return cnf, nil
}
