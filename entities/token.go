package entities

import (
	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	*jwt.Token
}
