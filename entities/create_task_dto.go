package entities 

import (
	"time"
)

type CreateTaskDto struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Deadline *time.Time `json:"deadline"`
	Tags []Tag `json:"tags,omitempty"`
}
