package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

// TODO поменять local на prod
type Config struct {
	Env string `env:"ENV" env-default:"local" env-required:"true"`
	Listen
	StorageConfig
}

// TODO Разобраться с таймаутами
type Listen struct {
	HostEndpoint string        `env:"HOST" env-default:"localhost"`
	PortEndpoint string        `env:"PORT" env-default:"8080"`
	WithEndpoint bool          `env:"WITH_ENDPOINT" env-default:"true"`
	Timeout      time.Duration `env:"TIMEOUT" env-default:"4s"`
	IdleTimeout  time.Duration `env:"IDLE_TIMEOUT" env-default:"60s"`
}

type StorageConfig struct {
	DBHost     string `env:"DB_HOST" env-default:"localhost"`
	DBPort     int    `env:"DB_PORT" env-default:"5432"`
	DBUsername string `env:"DB_USER" env-default:"postgres"`
	DBPassword string `env:"DB_PASSWORD" env-default:"password"`
	DBName     string `env:"DB_NAME" env-default:"wallet_db"`
	DBSSLMode  string `env:"DB_SSLMODE" env-default:"disable"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	// -c ="path/to/config.env"
	flag.StringVar(&res, "c", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
