package config

import "github.com/caarlos0/env/v9"

type Config struct {
	DBHost     string `env:"MYSQL_HOST" envDefault:"127.0.0.1"`
	DBUser     string `env:"MYSQL_USER" envDefault:"user"`
	DBPassword string `env:"MYSQL_PASSWORD" envDefault:"password"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
