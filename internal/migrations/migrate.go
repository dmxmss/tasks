package main

import (
	"github.com/dmxmss/tasks/config"
	"github.com/dmxmss/tasks/entities"
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

	migrate(db)
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Task{})
}
