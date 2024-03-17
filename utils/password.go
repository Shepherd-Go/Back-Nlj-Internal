package utils

import (
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

type Password interface {
	HashPassword(password *string)
	CheckPasswordHash(password, hash string) bool
	GenerateTemporaryPassword() string
}

type password struct{}

func NewHashPassword() Password {
	return &password{}
}

func (p *password) HashPassword(password *string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	*password = string(hash)
}

func (p *password) CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func (p *password) GenerateTemporaryPassword() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@#$%&"
	password := make([]byte, 8)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}
	return string(password)
}
