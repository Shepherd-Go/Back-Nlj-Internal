package config

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
