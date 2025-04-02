package entities

import (
	"time"
)

type GetTaskDto struct {
	Name string
	Description *string 
	Status *string
	Deadline *time.Time
	Tags []Tag 
	Weather *string
	UserID uint
}
