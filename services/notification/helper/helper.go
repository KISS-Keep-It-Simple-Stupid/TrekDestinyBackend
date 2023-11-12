package helper

import (
	"encoding/json"
	"net/http"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/notification/models"
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
func MessageGenerator(w http.ResponseWriter, respText string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errorMsg := struct {
		Message string `json:"message"`
	}{
		Message: respText,
	}

	errorResp, _ := json.Marshal(&errorMsg)
	w.Write(errorResp)
}

func ResponseGenerator(w http.ResponseWriter, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp, _ := json.Marshal(response)
	w.Write(resp)
}