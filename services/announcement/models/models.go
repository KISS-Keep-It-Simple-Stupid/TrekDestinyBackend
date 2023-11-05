package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	jwt.RegisteredClaims
	UserName string
}

type NotificationMessage struct {
	UserID  int    `json:"user_id"`
	Message string `json:"message"`
}
