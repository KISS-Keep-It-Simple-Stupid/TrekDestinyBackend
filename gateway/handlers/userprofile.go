package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/gateway/helpers"
	userprofile_pb "github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/gateway/services/userprofile"
)

func (s *Repository) Profile(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Jwt ")
	if len(splitToken) < 2 {
		helpers.MessageGenerator(w, "User is UnAuthorized", http.StatusUnauthorized)
		return
	}
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

func (s *Repository) EditProfile(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Jwt ")
	if len(splitToken) < 2 {
		helpers.MessageGenerator(w, "User is UnAuthorized", http.StatusUnauthorized)
		return
	}
	reqToken = splitToken[1]
	editReq := &userprofile_pb.EditProfileRequest{}
	postData, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body format", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(postData, editReq)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body fields", http.StatusBadRequest)
		return
	}
	editReq.AccessToken = reqToken
	resp, err := s.userprofile_client.EditProfile(context.Background(), editReq)
	if err != nil {
		helpers.MessageGenerator(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if resp.Message == "success" {
		helpers.MessageGenerator(w, "profile edited successfully", http.StatusOK)
	} else {
		helpers.MessageGenerator(w, resp.Message, http.StatusBadRequest)
	}
}
