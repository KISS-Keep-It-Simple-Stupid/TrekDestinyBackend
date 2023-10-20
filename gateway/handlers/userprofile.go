package handlers

import (
	"context"
	"net/http"
	"strings"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/gateway/helpers"
	userprofile_pb "github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/gateway/services/userprofile"
)

func (s *Repository) Profile(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Jwt ")
	reqToken = splitToken[1]
	profileReq := &userprofile_pb.ProfileDetailsRequest{AccessToken: reqToken}
	resp, err := s.userprofile_client.ProfileDetails(context.Background(), profileReq)
	if err != nil {
		helpers.MessageGenerator(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if resp.Message == "success" {
		helpers.ResponseGenerator(w, resp)
	} else {
		helpers.MessageGenerator(w, resp.Message, http.StatusUnauthorized)
	}
}
