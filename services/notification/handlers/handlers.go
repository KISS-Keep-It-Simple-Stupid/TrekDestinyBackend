package handlers

import (
	"net/http"
	"strings"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/notification/helper"
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
