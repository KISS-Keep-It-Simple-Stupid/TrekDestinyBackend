package helper

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/notification/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func TestMessageGenerator(t *testing.T) {
	recorder := httptest.NewRecorder()
	MessageGenerator(recorder, "this is test response", http.StatusOK)
}

func TestResponseGenerator(t *testing.T) {
	recorder := httptest.NewRecorder()
	ResponseGenerator(recorder, struct{ Message string }{Message: "this is test response"})
}

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