package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	jwt.RegisteredClaims
	UserName string
	UserID   int
}

type LoginCridentials struct {
	ID         int
	Email      string
	Password   string
	UserName   string
	FirstName  string
	IsVerified bool
}
