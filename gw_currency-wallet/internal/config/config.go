package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

// TODO поменять local на prod
type Config struct {
	Env string `env:"ENV" env-default:"local" env-required:"true"`
	Listen
}

// TODO Разобраться с таймаутами
type Listen struct {
	HostEndpoint string        `env:"HOST" env-default:"localhost"`
	PortEndpoint int           `env:"PORT" env-default:"8080"`
	WithEndpoint bool          `env:"WITH_ENDPOINT" env-default:"true"`
	Timeout      time.Duration `env:"TIMEOUT" env-default:"4s"`
	IdleTimeout  time.Duration `env:"IDLE_TIMEOUT" env-default:"60s"`
}

func MustLoad(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot open config: %s", err)
	}

	return &cfg
}
