package entities

import (
	"time"
)

type Task struct {
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Description *string 
	Status *string
	Deadline *time.Time
	Tags []Tag `gorm:"many2many:task_tags"`
	Weather *string
	UserID uint
	User User
}
