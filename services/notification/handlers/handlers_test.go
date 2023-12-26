package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/notification/models"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

type MockDBRepository struct {
}

func (m *MockDBRepository) InsertNotification(message models.NotifMessage) error {
	return nil
}

func (m *MockDBRepository) GetNotificationByID(userid int) ([]*models.NotifResponse, error) {
	return make([]*models.NotifResponse, 0), nil
}
func (m *MockDBRepository) DeleteNotification(notifID int) error {
	return nil
}


type MockDBRepository2 struct {
}

func (m *MockDBRepository2) InsertNotification(message models.NotifMessage) error {
	return nil
}

func (m *MockDBRepository2) GetNotificationByID(userid int) ([]*models.NotifResponse, error) {
	return make([]*models.NotifResponse, 0), nil
}
func (m *MockDBRepository2) DeleteNotification(notifID int) error {
	return errors.New("fake error")
}
func TestGetNotif(t *testing.T) {
	mockDB := &MockDBRepository{}
	repo := &Repository{DB: mockDB}
	viper.Set("JWTKEY", "as you wish")
	req, err := http.NewRequest("GET", "/notif", nil)
	assert.NoError(t, err)
	req.Header.Set("Authorization", "Jwt eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQzNjM5NDksIlVzZXJOYW1lIjoic29iaGFua2F6ZW1pIiwiVXNlcklEIjoxfQ.UblUPLZgHa25ztlzjyYT-vrhem7D2acIr2HEpFF_wEg")
	w := httptest.NewRecorder()
	repo.GetNotif(w, req)
	req2, err := http.NewRequest("GET", "/notif", nil)
	assert.NoError(t, err)
	req2.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQzNjM5NDksIlVzZXJOYW1lIjoic29iaGFua2F6ZW1pIiwiVXNlcklEIjoxfQ.UblUPLZgHa25ztlzjyYT-vrhem7D2acIr2HEpFF_wEg")
	w2 := httptest.NewRecorder()
	repo.GetNotif(w2, req2)
	req3, err := http.NewRequest("GET", "/notif", nil)
	assert.NoError(t, err)
	req3.Header.Set("Authorization", "Jwt asd")
	w3 := httptest.NewRecorder()
	repo.GetNotif(w3, req3)
}

func TestDeleteNotif(t *testing.T) {
	mockDB := &MockDBRepository{}
	viper.Set("JWTKEY", "as you wish")
	repo := &Repository{DB: mockDB}

	deleteNotif := models.DeleteNotifMessage{
		NotifID: 1,
	}
	reqBody, err := json.Marshal(deleteNotif)
	assert.NoError(t, err)

	req, err := http.NewRequest("DELETE", "/notif", bytes.NewBuffer(reqBody))
	assert.NoError(t, err)
	req.Header.Set("Authorization", "Jwt eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQzNjM5NDksIlVzZXJOYW1lIjoic29iaGFua2F6ZW1pIiwiVXNlcklEIjoxfQ.UblUPLZgHa25ztlzjyYT-vrhem7D2acIr2HEpFF_wEg")
	w := httptest.NewRecorder()
	repo.DeleteNotif(w, req)
	req2, err := http.NewRequest("DELETE", "/notif", bytes.NewBuffer(reqBody))
	assert.NoError(t, err)
	req2.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQzNjM5NDksIlVzZXJOYW1lIjoic29iaGFua2F6ZW1pIiwiVXNlcklEIjoxfQ.UblUPLZgHa25ztlzjyYT-vrhem7D2acIr2HEpFF_wEg")
	w2 := httptest.NewRecorder()
	repo.DeleteNotif(w2, req2)

	req3, err := http.NewRequest("DELETE", "/notif", bytes.NewBuffer(reqBody))
	assert.NoError(t, err)
	req3.Header.Set("Authorization", "Jwt JhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQzNjM5NDksIlVzZXJOYW1lIjoic29iaGFua2F6ZW1pIiwiVXNlcklEIjoxfQ.UblUPLZgHa25ztlzjyYT-vrhem7D2acIr2HEpFF_wEg")
	w3 := httptest.NewRecorder()
	repo.DeleteNotif(w3, req3)




	mockDB2 := &MockDBRepository2{}
	repo2 := &Repository{DB: mockDB2}
	req4, err := http.NewRequest("DELETE", "/notif", bytes.NewBuffer(reqBody))
	assert.NoError(t, err)
	req4.Header.Set("Authorization", "Jwt eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQzNjM5NDksIlVzZXJOYW1lIjoic29iaGFua2F6ZW1pIiwiVXNlcklEIjoxfQ.UblUPLZgHa25ztlzjyYT-vrhem7D2acIr2HEpFF_wEg")
	w4 := httptest.NewRecorder()
	repo2.DeleteNotif(w4, req4)

}
