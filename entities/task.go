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
	Tags []string
	Weather *string
	UserID uint
	User User
}
