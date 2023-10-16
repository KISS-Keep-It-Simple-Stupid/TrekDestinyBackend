package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	jwt.RegisteredClaims
	UserName string
	ExpDate  time.Time
}

type LoginCridentials struct {
	Email      string
	Password   string
	UserName   string
	FirstName  string
	IsVerified bool
}
