package entities

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID int `json:"user_id"`
	City string `json:"city,omitempty"`
	jwt.RegisteredClaims
}  
