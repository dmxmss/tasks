package entities

import (
	"time"
)

type GetTaskDto struct {
	Name string `json:"name"`
	Description *string `json:"description;omitempty"`
	Status *string `json:"status;omitempty"`
	Deadline *time.Time `json:"deadline;omitempty"`
	Weather *string `json:"weather;omitempty"`
	UserID uint `json:"user_id"`
}
