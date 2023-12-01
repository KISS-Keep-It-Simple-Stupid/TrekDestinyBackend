package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/chat/models"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

type MockDBRepository struct {
}

func (m *MockDBRepository) GetCount(chatID string) int {
	return 10
}

func (m *MockDBRepository) InsertChatMessage(message *models.Message) {

}
func (m *MockDBRepository) GetHistory(chat_id string, PageNumber, numberOfMessages int) ([]models.ChatMessage, error) {
	return []models.ChatMessage{}, nil
}

type MockDBRepository2 struct {
}

func (m *MockDBRepository2) GetCount(chatID string) int {
	return 0
}

func (m *MockDBRepository2) InsertChatMessage(message *models.Message) {

}
func (m *MockDBRepository2) GetHistory(chat_id string, PageNumber, numberOfMessages int) ([]models.ChatMessage, error) {
	return []models.ChatMessage{}, errors.New("test error")
}

func TestNewHandler(t *testing.T) {
	mockHub := &models.Hub{}
	NewHandler(mockHub, &MockDBRepository{})
}
func TestChat(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/count/123/456", nil)
	assert.NoError(t, err, "Error creating HTTP request")
	handler := &Repository{
		Hub: &models.Hub{},
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
	}

	handler.Chat(recorder, request)

}

func TestCount(t *testing.T) {
	responseRecorder := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/count/123/456", nil)
	assert.NoError(t, err, "Error creating HTTP request")
	mockDBRepository := &MockDBRepository{}
	repo := &Repository{
		DB: mockDBRepository,
	}
	repo.Count(responseRecorder, request)
}

func TestHistory(t *testing.T) {
	responseRecorder := httptest.NewRecorder()
	request1, _ := http.NewRequest("GET", "/history/123/456?page=1&count=2", nil)
	request4, _ := http.NewRequest("GET", "/history/123/456?page=1&count=2", nil)
	request2, _ := http.NewRequest("GET", "/history/123/456?page=1&count=asda", nil)
	request3, _ := http.NewRequest("GET", "/history/123/456?page=asd&count=2", nil)
	reqs := make([]*http.Request, 4)
	reqs[0] = request1
	reqs[1] = request2
	reqs[2] = request3
	reqs[3] = request4
	repo := &Repository{}
	for i := 0; i < 4; i++ {
		if i == 0 {
			repo.DB = &MockDBRepository2{}
		} else {
			repo.DB = &MockDBRepository{}
		}
		repo.History(responseRecorder, reqs[i])
	}
}
