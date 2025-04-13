package entities

import (
	"time"
)

type GetTaskDto struct {
	ID int `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Status *string `json:"status,omitempty"`
	Deadline *time.Time `json:"deadline,omitempty"`
	Tags []Tag `json:"tags,omitempty"`
	Weather *string `json:"weather,omitempty"`
	UserID int `json:"user_id"`
}
