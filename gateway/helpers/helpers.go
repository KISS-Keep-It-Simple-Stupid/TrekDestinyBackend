package helpers

import (
	"encoding/json"
	"net/http"
)

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
