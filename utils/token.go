package utils

import (
	"github.com/form3tech-oss/jwt-go"
)

func GetToken() *jwt.Token {
	return jwt.New(jwt.SigningMethodHS256)
}

