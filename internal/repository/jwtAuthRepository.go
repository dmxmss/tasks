package repository

import (
	"github.com/dmxmss/tasks/config"
	"github.com/dmxmss/tasks/entities"
	e "github.com/dmxmss/tasks/error"
	"github.com/golang-jwt/jwt/v5"

	"errors"
	"time"
)

type AuthRepository interface {
	ValidateToken(string) (*jwt.Token, error)
	GenerateAccessToken(int) *jwt.Token
	GenerateRefreshToken() *jwt.Token
}

type jwtAuthRepository struct {
	conf *config.Auth
}

func NewAuthRepository(conf *config.Auth) AuthRepository {
	return &jwtAuthRepository{
		conf: conf,
	}
}

func (jwtRepo *jwtAuthRepository) ValidateToken(rawToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(rawToken, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwtRepo.conf.SigningMethod {
			return nil, jwt.ErrSignatureInvalid
		}

		return jwtRepo.conf.JWTSecret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, e.ErrAuthSignatureInvalid
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, e.ErrAuthTokenExpired
		} else {
			return nil, e.ErrAuthFailed
		}
	}

	return token, nil
}

func (jwtRepo *jwtAuthRepository) GenerateAccessToken(userId int) *jwt.Token {
	exp := time.Now().Add(time.Duration(jwtRepo.conf.Access.ExpirationTime)*time.Second)

	claims := entities.Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwtRepo.conf.SigningMethod, claims)
	return token
}

func (jwtRepo *jwtAuthRepository) GenerateRefreshToken() *jwt.Token {
	exp := time.Now().Add(time.Duration(jwtRepo.conf.Refresh.ExpirationTime)*time.Second)

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(exp),
	}

	token := jwt.NewWithClaims(jwtRepo.conf.SigningMethod, claims)
	return token
}
