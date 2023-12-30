package helper

import (
	"testing"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/userprofile/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func TestDecodeToken(t *testing.T) {
	// Set up the key for testing (replace with your actual key)
	viper.Set("JWTKEY", "your_test_key")
	defer viper.Reset()

	// Create a sample JWT token
	claims := &models.JwtClaims{}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(viper.GetString("JWTKEY")))
	if err != nil {
		t.Fatalf("Failed to sign JWT token: %v", err)
	}
	DecodeToken(signedToken)
}
