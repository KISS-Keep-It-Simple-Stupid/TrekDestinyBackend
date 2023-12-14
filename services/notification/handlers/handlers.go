package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/notification/helper"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/notification/models"
)

func (s *Repository) GetNotif(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Jwt ")
	if len(splitToken) < 2 {
		helper.MessageGenerator(w, "User is UnAuthorized", http.StatusUnauthorized)
		return
	}
	reqToken = splitToken[1]
	claims, err := helper.DecodeToken(reqToken)
	if err != nil {
		helper.MessageGenerator(w, "User is UnAuthorized", http.StatusUnauthorized)
		return
	}
	notifs, err := s.DB.GetNotificationByID(claims.UserID)
	if err != nil {
		helper.MessageGenerator(w, "internal error while getting notifcations - notification service", http.StatusInternalServerError)
		return
	}
	helper.ResponseGenerator(w, notifs)
}

func (s *Repository) DeleteNotif(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Jwt ")
	if len(splitToken) < 2 {
		helper.MessageGenerator(w, "User is UnAuthorized", http.StatusUnauthorized)
		return
	}
	reqToken = splitToken[1]
	_, err := helper.DecodeToken(reqToken)
	if err != nil {
		helper.MessageGenerator(w, "User is UnAuthorized", http.StatusUnauthorized)
		return
	}

	deleteNotification_req := models.DeleteNotifMessage{}
	postData, err := io.ReadAll(r.Body)
	if err != nil {
		helper.MessageGenerator(w, "wrong post body format", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(postData, &deleteNotification_req)
	if err != nil {
		helper.MessageGenerator(w, "wrong post body fields", http.StatusBadRequest)
		return
	}
	err = s.DB.DeleteNotification(deleteNotification_req.NotifID)
	if err != nil {
		helper.MessageGenerator(w, "internal error while deleting a notification - notification service", http.StatusInternalServerError)
		return
	}
	helper.MessageGenerator(w, "Notification Deleted", http.StatusOK)
}
