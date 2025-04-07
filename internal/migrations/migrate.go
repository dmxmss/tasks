package main

import (
	"github.com/dmxmss/tasks/config"
	"github.com/dmxmss/tasks/entities"
	"github.com/dmxmss/tasks/internal"
	"gorm.io/gorm"
)

func main() {
	conf := config.GetConfig()

	tasksRepo, err := internal.NewPgTasksRepository(conf)
	if err != nil {
		panic(err)
	}

	migrate(tasksRepo.GetDb())
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Task{})
}
