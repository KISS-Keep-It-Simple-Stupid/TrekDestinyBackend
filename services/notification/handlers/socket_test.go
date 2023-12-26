package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"


	"github.com/DATA-DOG/go-sqlmock"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/notification/dbrepo"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/notification/models"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestSubscribe(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/1", nil)
	assert.NoError(t, err, "Error creating HTTP request")
	handler := &Repository{
		hub: &models.Hub{},
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
	}

	handler.Subscribe(recorder, request)
}

func TestNew(t *testing.T) {
	mockDB, _, _ := sqlmock.New()
	repo := dbrepo.New(mockDB)
	hub := &models.Hub{}
	New(hub , repo)
}

