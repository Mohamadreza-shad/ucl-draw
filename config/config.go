package config

import (
	"flag"

	"github.com/pkg/errors"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/source/env"
)

const (
	EnvProd = "prod"
	EnvTest = "test"
	EnvDev  = "dev"
)

type Config struct {
	Redis    Redis
	Env      string
	Postgres Postgres
}

var cfg *Config = &Config{}

func GetEnv() string {
	env := cfg.Env
	if env != "" {
		return env
	}
	return EnvProd
}

func SetTestEnvVariable() {
	cfg.Env = EnvTest
}

func IsTestEnv() bool {
	return flag.Lookup("test.v") != nil
}

func Load() error {
	microConfig, err := config.NewConfig(config.WithSource(env.NewSource()))
	if err != nil {
		return errors.Wrap(err, "config.New")
	}
	if err := microConfig.Load(); err != nil {
		return errors.Wrap(err, "config.Load")
	}
	if err := microConfig.Scan(cfg); err != nil {
		return errors.Wrap(err, "config.Scan")
	}
	return nil
}
