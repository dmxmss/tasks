package entities

import (
	"time"
)

type PatchTaskDto struct {
	ID int `json:"id"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Status *string `json:"status,omitempty"`
	Deadline *time.Time `json:"deadline,omitempty"`
	Tags []Tag `gorm:"many2many:task_tags" json:"tags,omitempty"`
	Weather *string `json:"weather,omitempty"`
}
