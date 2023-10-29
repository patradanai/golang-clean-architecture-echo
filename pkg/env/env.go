package cfg

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type (
	envConfigs struct {
		Http     http
		Database database
	}

	database struct {
		Uri string `env:"MONGO_URI" envDefault:"mongodb://localhost:27017"`
	}

	http struct {
		Port string `env:"PORT" envDefault:"8080"`
	}
)

var cfg envConfigs

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}

}

func Get() envConfigs {
	return cfg
}
