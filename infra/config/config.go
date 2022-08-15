package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DatabaseUser     string `envconfig:"DATABASE_USER" default:"root"`
	DatabasePassword string `envconfig:"DATABASE_PASSWORD" default:"root"`
	DatabaseHost     string `envconfig:"DATABASE_HOST" default:"localhost"`
	DatabaseName     string `envconfig:"DATABASE_NAME" default:"unico"`
	DatabasePort     string `envconfig:"DATABASE_PORT" default:"3306"`
	Port             string `envconfig:"PORT" default:"8000"`
	PathCSV          string `envconfig:"PATH_CVS" default:"DEINFO_AB_FEIRASLIVRES_2014.csv"`
}

func SetupEnvFile() *Config {
	envConfig := &Config{}
	_ = godotenv.Load()
	err := envconfig.Process("", envConfig)
	if err != nil {
		log.Fatal(nil, "Fatal error ", err)
	}

	return envConfig
}
