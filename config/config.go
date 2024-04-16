package config

// mockery --all --disable-version-string --keeptree ----> comando para generar mocks automaticos...

import (
	"log"
	"os"
	"sync"

	"github.com/andresxlp/gosuite/config"
)

var (
	Once sync.Once
	Cfg  Config
)

type Config struct {
	Server   Server   `validate:"required" mapstructure:"server"`
	JWT      JWT      `validate:"required" mapstructure:"jwt"`
	Database Database `validate:"required" mapstructure:"database"`
	Email    Email    `validate:"required" mapstructure:"email"`
}

type Server struct {
	Port int `validate:"required" mapstructure:"port"`
}

type JWT struct {
	Key string `validate:"required" mapstructure:"key"`
}

type Database struct {
	Host     string `validate:"required" mapstructure:"host"`
	Port     int    `validate:"required" mapstructure:"port"`
	User     string `validate:"required" mapstructure:"user"`
	Password string `validate:"required" mapstructure:"password"`
	DBName   string `validate:"required" mapstructure:"dbname"`
}

type Email struct {
	Email    string `validate:"required" mapstructure:"email"`
	Password string `validate:"required" mapstructure:"password"`
}

func Environments() Config {
	Once.Do(func() {
		if os.Getenv("DEVMODE") == "true" {

			if err := config.SetEnvsFromFile(".env"); err != nil {
				log.Panic(err)
			}
		}

		if err := config.GetConfigFromEnv(&Cfg); err != nil {
			log.Panicf("Error partsing enviroment vars %v", err)
		}

	})

	return Cfg
}
