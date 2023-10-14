package helper

import (
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/authentication/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func NewToken(claims *models.JwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := viper.Get("JWTKEY").(string)
	token_string, err := token.SignedString([]byte(key))
	if err != nil {
		return "", nil
	}
	return token_string, nil
}
