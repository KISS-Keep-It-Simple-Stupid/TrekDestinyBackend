package helpers

import (
	"encoding/json"
	"net/http"
)


func RespGenerator(w http.ResponseWriter, respText string, statusCode int) {
	w.Header().Set("Content-Type" , "application/json")
	w.WriteHeader(statusCode)
	errorMsg := struct {
		Message string `json:"message"`
	}{
		Message: respText,
	}

	errorResp, _ := json.Marshal(&errorMsg)
	w.Write(errorResp)
}