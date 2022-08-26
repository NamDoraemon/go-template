package configs

import (
	"github.com/caarlos0/env/v6"
	"log"
)

var GlobalConf Configs

type Configs struct {
	AppEnv   string `env:"APP_ENV"`
	HTTPPort string `env:"HTTP_PORT" envDefault:"6001"`
	GRPCPort string `env:"GRPC_PORT" envDefault:"6000"`

	DBUser string `env:"DB_USER"`
	DBPass string `env:"DB_PASS"`
	DBHost string `env:"DB_HOST"`
	DBPort string `env:"DB_PORT""`
	DBName string `env:"DB_NAME"`

	RedisHost string `env:"REDIS_HOST"`
	RedisPort string `env:"REDIS_PORT"`
	RedisDB   int    `env:"REDIS_DB"`
}

func Load() {
	if err := env.Parse(&GlobalConf); err != nil {
		log.Fatal(err)
	}
}
