package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/gateway/helpers"
	announcement_pb "github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/gateway/services/announcement"
)



func (s *Repository) CreateCard(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Jwt ")
	if len(splitToken) < 2 {
		helpers.MessageGenerator(w, "User is UnAuthorized", http.StatusUnauthorized)
		return
	}
	reqToken = splitToken[1]
	createcardReq := &announcement_pb.CreateCardRequest{}
	postData, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body format", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(postData, createcardReq)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body fields", http.StatusBadRequest)
		return
	}
	createcardReq.AccessToken = reqToken
	resp, err := s.announcement_client.CreateCard(context.Background(), createcardReq)
	if err != nil {
		helpers.MessageGenerator(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if resp.Message == "success" {
		helpers.MessageGenerator(w, "announcement created successfully", http.StatusCreated)
	} else if resp.Message == "User is UnAuthorized - announcement service" {
		helpers.MessageGenerator(w, resp.Message, http.StatusUnauthorized)
	} else {
		helpers.MessageGenerator(w, resp.Message, http.StatusBadRequest)
	}
}

func (s *Repository) GetCard(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Jwt ")
	if len(splitToken) < 2 {
		helpers.MessageGenerator(w, "User is UnAuthorized", http.StatusUnauthorized)
		return
	}
	reqToken = splitToken[1]
	getcardReq := &announcement_pb.GetCardRequest{}
	getcardReq.AccessToken = reqToken

	queryParams := r.URL.Query()
	filterValues := queryParams.Get("filter")
	sortValue := queryParams.Get("sort")
	pageSize := queryParams.Get("page-size")
	pageNumber := queryParams.Get("page-number")
	getcardReq.FilterValues = filterValues
	getcardReq.SortValue = sortValue
	getcardReq.PageSize = pageSize
	getcardReq.PageNumber = pageNumber

	resp, err := s.announcement_client.GetCard(context.Background(), getcardReq)
	if err != nil {
		helpers.MessageGenerator(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if resp.Message == "success" {
		helpers.ResponseGenerator(w, resp)
	} else if resp.Message == "User is UnAuthorized - announcement service" {
		helpers.MessageGenerator(w, resp.Message, http.StatusUnauthorized)
	} else {
		helpers.MessageGenerator(w, resp.Message, http.StatusBadRequest)
	}
}
