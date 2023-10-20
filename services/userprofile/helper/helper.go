package helper

import (
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/userprofile/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func DecodeToken(token string) (*models.JwtClaims, error) {
	claims := &models.JwtClaims{}
	key := viper.Get("JWTKEY").(string)
	jwttoken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}
	if !jwttoken.Valid {
		return nil, err
	}
	return claims, nil
}
