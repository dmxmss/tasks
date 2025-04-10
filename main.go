package main

import (
	"github.com/dmxmss/tasks/server"
	"github.com/dmxmss/tasks/config"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"

	"fmt"
)

func main() {
	conf := config.GetConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		conf.Database.Host,
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Name,
		conf.Database.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})	
	if err != nil {
		panic(err)
	}

	s, err := server.NewGinServer(conf, db)
	if err != nil {
		panic(err)
	}

	s.RegisterHandlers()

	s.Start()
}
