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
	ValidateToken(string) (*entities.Claims, error)
	GenerateTokens(int) (*string, *string, error)
}

type jwtAuthRepository struct {
	conf *config.Auth
}

func NewAuthRepository(conf *config.Auth) AuthRepository {
	return &jwtAuthRepository{
		conf: conf,
	}
}

func (jwtRepo *jwtAuthRepository) ValidateToken(rawToken string) (*entities.Claims, error) {
	token, err := jwt.ParseWithClaims(rawToken, &entities.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwtRepo.conf.SigningMethod {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(jwtRepo.conf.JWTSecret), nil
	})

	claims, ok := token.Claims.(*entities.Claims)

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, e.ErrAuthSignatureInvalid
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, e.ErrAuthTokenExpired
		} else {
			return nil, e.ErrAuthFailed
		}
	} else if !token.Valid || !ok {
		return nil, e.ErrAuthTokenInvalid
	}

	return claims, nil
}

func (jwtRepo *jwtAuthRepository) generateAccessToken(userId int) (*string, error) {
	exp := time.Now().Add(time.Duration(jwtRepo.conf.Access.ExpirationTime)*time.Second)

	claims := entities.Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	_, err := claims.GetExpirationTime()
	if err != nil {
		panic(err)
	}

	token := jwt.NewWithClaims(jwtRepo.conf.SigningMethod, claims)
	signedToken, err := token.SignedString([]byte(jwtRepo.conf.JWTSecret))
	if err != nil {
		return nil, e.ErrTokenSigningError
	}
	return &signedToken, nil
}

func (jwtRepo *jwtAuthRepository) generateRefreshToken() (*string, error) {
	exp := time.Now().Add(time.Duration(jwtRepo.conf.Refresh.ExpirationTime)*time.Second)

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(exp),
	}

	token := jwt.NewWithClaims(jwtRepo.conf.SigningMethod, claims)
	signedToken, err := token.SignedString([]byte(jwtRepo.conf.JWTSecret))
	if err != nil {
		return nil, e.ErrTokenSigningError
	}
	return &signedToken, nil
}

func (jwtRepo *jwtAuthRepository) GenerateTokens(userId int) (*string, *string, error) {
	accessToken, err := jwtRepo.generateAccessToken(userId)
	if err != nil {
		return nil, nil, err
	}
	refreshToken, err := jwtRepo.generateRefreshToken()
	if err != nil {
		return nil, nil, err
	}
	return accessToken, refreshToken, nil
}
