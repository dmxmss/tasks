package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string
	Email string
	Password string
	City string
}
