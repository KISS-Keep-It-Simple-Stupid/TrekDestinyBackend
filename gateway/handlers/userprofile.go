package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/gateway/helpers"
	announcement_pb "github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/gateway/services/announcement"
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

func (s *Repository) UploadImage(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Jwt ")
	if len(splitToken) < 2 {
		helpers.MessageGenerator(w, "User is UnAuthorized", http.StatusUnauthorized)
		return
	}
	reqToken = splitToken[1]
	isHostImage := r.URL.Query().Get("host")
	isBlogImage := r.URL.Query().Get("blog")
	// Parse the form data, including the uploaded file
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit for the image size
	if err != nil {
		helpers.MessageGenerator(w, "image size is too large", http.StatusBadRequest)
		return
	}
	if isHostImage != "" {
		count, _ := strconv.Atoi(r.URL.Query().Get("count"))
		hostImageUploadReq := &announcement_pb.HostHouseImageRequest{
			AccessToken: reqToken,
			ImageData:   make([][]byte, 0),
		}
		for i := 1; i <= count; i++ {
			file, _, err := r.FormFile(fmt.Sprintf("image%d", i))
			if err != nil {
				log.Println(err.Error())
				helpers.MessageGenerator(w, "inavlid key image", http.StatusBadRequest)
				return
			}
			defer file.Close()
			image, err := io.ReadAll(file)
			if err != nil {
				log.Println(err.Error())
				helpers.MessageGenerator(w, "corrupted image", http.StatusBadRequest)
				return
			}
			hostImageUploadReq.ImageData = append(hostImageUploadReq.ImageData, image)
		}

		resp, err := s.announcement_client.UploadHostHouseImage(context.Background(), hostImageUploadReq)
		if err != nil {
			helpers.MessageGenerator(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if resp.Message == "success" {
			helpers.MessageGenerator(w, "host images uploaded", http.StatusOK)
		} else {
			helpers.MessageGenerator(w, resp.Message, http.StatusBadRequest)
		}
		return
	}

	// Get the file from the form data
	file, _, err := r.FormFile("image")
	if err != nil {
		helpers.MessageGenerator(w, "There is no image field in form data", http.StatusBadRequest)
		return
	}
	defer file.Close()
	image_data, err := io.ReadAll(file)
	if err != nil {
		helpers.MessageGenerator(w, "Image file is corrupted", http.StatusBadRequest)
		return
	}
	if isBlogImage != "" {
		blogID, _ := strconv.Atoi(r.URL.Query().Get("id"))
		blogUploadReq := &announcement_pb.UploadBlogImageRequest{
			ImageData:   image_data,
			AccessToken: reqToken,
			BlogID:      int32(blogID),
		}
		resp, err := s.announcement_client.UploadBlogImage(context.Background(), blogUploadReq)
		if err != nil {
			helpers.MessageGenerator(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if resp.Message == "success" {
			helpers.MessageGenerator(w, "blog image uploaded", http.StatusOK)
		} else {
			helpers.MessageGenerator(w, resp.Message, http.StatusBadRequest)
		}
	} else {
		uploadReq := &userprofile_pb.ImageRequest{
			ImageData:   image_data,
			AccessToken: reqToken,
		}
		resp, err := s.userprofile_client.UploadImage(context.Background(), uploadReq)
		if err != nil {
			helpers.MessageGenerator(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if resp.Message == "success" {
			helpers.MessageGenerator(w, "profile image uploaded", http.StatusOK)
		} else {
			helpers.MessageGenerator(w, resp.Message, http.StatusBadRequest)
		}
	}
}

func (s *Repository) PublicProfile(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Jwt ")
	if len(splitToken) < 2 {
		helpers.MessageGenerator(w, "User is UnAuthorized", http.StatusUnauthorized)
		return
	}
	reqToken = splitToken[1]
	username := chi.URLParam(r, "username")
	profileReq := &userprofile_pb.PublicProfileRequest{
		AccessToken: reqToken,
		Username:    username,
	}
	resp, err := s.userprofile_client.PublicProfile(context.Background(), profileReq)
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

func (s *Repository) PublicProfileHost(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Jwt ")
	if len(splitToken) < 2 {
		helpers.MessageGenerator(w, "User is UnAuthorized", http.StatusUnauthorized)
		return
	}
	reqToken = splitToken[1]
	publicprofilehostReq := &userprofile_pb.PublicProfileHostRequest{}
	postData, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body format", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(postData, publicprofilehostReq)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body fields", http.StatusBadRequest)
		return
	}
	publicprofilehostReq.AccessToken = reqToken
	resp, err := s.userprofile_client.PublicProfileHost(context.Background(), publicprofilehostReq)
	if err != nil {
		helpers.MessageGenerator(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if resp.Message == "success" {
		helpers.ResponseGenerator(w, resp)
	} else if resp.Message == "you can't view this profile because the user hasn't offered to any of your announcements" {
		helpers.MessageGenerator(w, resp.Message, http.StatusBadRequest)
	} else {
		helpers.MessageGenerator(w, resp.Message, http.StatusUnauthorized)
	}
}

func (s *Repository) ChatListPost(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Jwt ")
	if len(splitToken) < 2 {
		helpers.MessageGenerator(w, "User is UnAuthorized", http.StatusUnauthorized)
		return
	}
	reqToken = splitToken[1]
	addChatListReq := &userprofile_pb.AddChatListRequest{}
	postData, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body format", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(postData, addChatListReq)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body fields", http.StatusBadRequest)
		return
	}
	addChatListReq.AccessToken = reqToken
	resp, err := s.userprofile_client.AddToChatList(context.Background(), addChatListReq)
	if err != nil {
		helpers.MessageGenerator(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if resp.Message == "success" {
		helpers.ResponseGenerator(w, resp)
	} else {
		helpers.MessageGenerator(w, resp.Message, http.StatusBadRequest)
	}
}

func (s *Repository) ChatList(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Jwt ")
	if len(splitToken) < 2 {
		helpers.MessageGenerator(w, "User is UnAuthorized", http.StatusUnauthorized)
		return
	}
	reqToken = splitToken[1]
	chatListReq := &userprofile_pb.ChatListRequest{AccessToken: reqToken}
	resp, err := s.userprofile_client.GetChatList(context.Background(), chatListReq)
	if err != nil {
		helpers.MessageGenerator(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if resp.Message == "success" {
		helpers.ResponseGenerator(w, resp)
	} else {
		helpers.MessageGenerator(w, resp.Message, http.StatusBadRequest)
	}
}
