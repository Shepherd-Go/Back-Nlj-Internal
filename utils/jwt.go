package utils

import (
	"time"

	"github.com/Shepherd-Go/Back-Nlj-Internal.git/config"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/dtos"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	key = config.Environments().JWT.Key
)

type JWT interface {
	SignedLoginToken(u dtos.EmployeeResponse) (string, error)
	PaserLoginJWT(value string) (jwt.MapClaims, error)
}

type jwtUtils struct{}

func NewJWT() JWT {
	return &jwtUtils{}
}

type MyCustomClaims struct {
	ID          uuid.UUID `json:"id"`
	Permissions string    `json:"permissions"`
	jwt.RegisteredClaims
}

func (jwtutils *jwtUtils) SignedLoginToken(e dtos.EmployeeResponse) (string, error) {

	claims := MyCustomClaims{
		e.ID,
		e.Permissions,
		jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour))},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(key))

}

func (jwtutils *jwtUtils) PaserLoginJWT(value string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(value, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}

	return token.Claims.(jwt.MapClaims), nil

}
