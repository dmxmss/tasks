package entities

import (
	"time"
)

type SearchTasksParams struct {
	Status *string `json:"status,omitempty"`
	Deadline *time.Time `json:"deadline,omitempty"`
}
