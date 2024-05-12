package db

import (
	"fmt"
	"log"

	"github.com/Shepherd-Go/Back-Nlj-Internal.git/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type pgOptions struct {
	host     string
	port     int
	user     string
	password string
	dbName   string
}

func (p *pgOptions) getDns() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.host, p.port, p.user, p.password, p.dbName)
}

func NewPostgresConnection() *gorm.DB {
	dns := pgOptions{
		host:     config.Environments().DatabasePG.Host,
		port:     config.Environments().DatabasePG.Port,
		user:     config.Environments().DatabasePG.User,
		password: config.Environments().DatabasePG.Password,
		dbName:   config.Environments().DatabasePG.DBName,
	}

	dbInstance, err := gorm.Open(postgres.Open(dns.getDns()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	log.Print("Postgres Connection Successfully")
	return dbInstance
}
