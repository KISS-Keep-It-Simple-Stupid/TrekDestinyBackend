package helper

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMessageGenerator(t *testing.T) {
	recorder := httptest.NewRecorder()
	MessageGenerator(recorder, "this is test response", http.StatusOK)
}

func TestResponseGenerator(t *testing.T) {
	recorder := httptest.NewRecorder()
	ResponseGenerator(recorder, struct{ Message string }{Message: "this is test response"})
}
