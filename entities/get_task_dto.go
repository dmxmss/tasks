package entities

import (
	"time"
)

type GetTaskDto struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Status *string `json:"status,omitempty"`
	Deadline *time.Time `json:"deadline,omitempty"`
	Tags []Tag `json:"tags,omitempty"`
	Weather *string `json:"weather,omitempty"`
	UserID uint `json:"user_id"`
}
