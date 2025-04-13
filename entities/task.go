package entities

import (
	"time"
)

type Task struct {
	ID int `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Name string `gorm:"not null"`
	Description *string 
	Status *string `gorm:"default:started"`
	Deadline *time.Time
	Tags []Tag `gorm:"many2many:task_tags"`
	Weather *string
	UserID int
	User User `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}
