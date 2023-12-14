package handlers

import (
	"context"
	"encoding/json"
	"fmt"
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

func (s *Repository) CreateOffer(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Jwt ")
	if len(splitToken) < 2 {
		helpers.MessageGenerator(w, "User is UnAuthorized", http.StatusUnauthorized)
		return
	}
	reqToken = splitToken[1]
	createofferReq := &announcement_pb.CreateOfferRequest{}
	postData, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body format", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(postData, createofferReq)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body fields", http.StatusBadRequest)
		return
	}
	createofferReq.AccessToken = reqToken
	resp, err := s.announcement_client.CreateOffer(context.Background(), createofferReq)
	if err != nil {
		helpers.MessageGenerator(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if resp.Message == "success" {
		helpers.MessageGenerator(w, "offer created successfully", http.StatusCreated)
	} else if resp.Message == "User is UnAuthorized - announcement service" {
		helpers.MessageGenerator(w, resp.Message, http.StatusUnauthorized)
	} else {
		helpers.MessageGenerator(w, resp.Message, http.StatusBadRequest)
	}
}

func (s *Repository) GetOffer(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Jwt ")
	if len(splitToken) < 2 {
		helpers.MessageGenerator(w, "User is UnAuthorized", http.StatusUnauthorized)
		return
	}
	reqToken = splitToken[1]
	getofferReq := &announcement_pb.GetOfferRequest{}
	postData, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body format", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(postData, getofferReq)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body fields", http.StatusBadRequest)
		return
	}
	getofferReq.AccessToken = reqToken
	resp, err := s.announcement_client.GetOffer(context.Background(), getofferReq)
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

func (s *Repository) GetCardProfile(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Jwt ")
	if len(splitToken) < 2 {
		helpers.MessageGenerator(w, "User is UnAuthorized", http.StatusUnauthorized)
		return
	}
	reqToken = splitToken[1]
	getcardprofileReq := &announcement_pb.GetCardProfileRequest{}
	getcardprofileReq.AccessToken = reqToken
	resp, err := s.announcement_client.GetCardProfile(context.Background(), getcardprofileReq)
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

func (s *Repository) CreatePost(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Jwt ")
	if len(splitToken) < 2 {
		helpers.MessageGenerator(w, "User is UnAuthorized", http.StatusUnauthorized)
		return
	}
	reqToken = splitToken[1]
	createpostReq := &announcement_pb.CreatePostRequest{}
	postData, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body format", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(postData, createpostReq)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body fields", http.StatusBadRequest)
		return
	}
	createpostReq.AccessToken = reqToken
	resp, err := s.announcement_client.CreatePost(context.Background(), createpostReq)
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

func (s *Repository) GetMyPost(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Jwt ")
	if len(splitToken) < 2 {
		helpers.MessageGenerator(w, "User is UnAuthorized", http.StatusUnauthorized)
		return
	}
	reqToken = splitToken[1]
	getmypostReq := &announcement_pb.GetMyPostRequest{}
	getmypostReq.AccessToken = reqToken
	resp, err := s.announcement_client.GetMyPost(context.Background(), getmypostReq)
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

func (s *Repository) GetPostHost(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Jwt ")
	if len(splitToken) < 2 {
		helpers.MessageGenerator(w, "User is UnAuthorized", http.StatusUnauthorized)
		return
	}
	reqToken = splitToken[1]
	getposthostReq := &announcement_pb.GetPostHostRequest{}
	postData, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body format", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(postData, getposthostReq)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body fields", http.StatusBadRequest)
		return
	}
	getposthostReq.AccessToken = reqToken
	resp, err := s.announcement_client.GetPostHost(context.Background(), getposthostReq)
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

func (s *Repository) AcceptOffer(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Jwt ")
	if len(splitToken) < 2 {
		helpers.MessageGenerator(w, "User is UnAuthorized", http.StatusUnauthorized)
		return
	}
	reqToken = splitToken[1]
	acceptofferReq := &announcement_pb.AcceptOfferRequest{}
	postData, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body format", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(postData, acceptofferReq)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body fields", http.StatusBadRequest)
		return
	}
	acceptofferReq.AccessToken = reqToken
	resp, err := s.announcement_client.AcceptOffer(context.Background(), acceptofferReq)
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

func (s *Repository) RejectOffer(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Jwt ")
	if len(splitToken) < 2 {
		helpers.MessageGenerator(w, "User is UnAuthorized", http.StatusUnauthorized)
		return
	}
	reqToken = splitToken[1]
	rejectofferReq := &announcement_pb.RejectOfferRequest{}
	postData, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body format", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(postData, rejectofferReq)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body fields", http.StatusBadRequest)
		return
	}
	rejectofferReq.AccessToken = reqToken
	resp, err := s.announcement_client.RejectOffer(context.Background(), rejectofferReq)
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

func (s *Repository) UploadPostImage(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Jwt ")
	if len(splitToken) < 2 {
		helpers.MessageGenerator(w, "User is UnAuthorized", http.StatusUnauthorized)
		return
	}
	reqToken = splitToken[1]
	// Parse the form data, including the uploaded file
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit for the image size
	if err != nil {
		helpers.MessageGenerator(w, "image size is too large", http.StatusBadRequest)
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
	uploadReq := &announcement_pb.PostImageRequest{
		AccessToken: reqToken,
		ImageData:   image_data,
	}
	resp, err := s.announcement_client.UploadPostImage(context.Background(), uploadReq)
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

func (s *Repository) EditAnnouncement(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Jwt ")
	if len(splitToken) < 2 {
		helpers.MessageGenerator(w, "User is UnAuthorized", http.StatusUnauthorized)
		return
	}
	reqToken = splitToken[1]
	editReq := &announcement_pb.EditAnnouncementRequest{}
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
	resp, err := s.announcement_client.EditAnnouncement(context.Background(), editReq)
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

func (s *Repository) DeleteAnnouncement(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Jwt ")
	if len(splitToken) < 2 {
		helpers.MessageGenerator(w, "User is UnAuthorized", http.StatusUnauthorized)
		return
	}
	reqToken = splitToken[1]
	deleteReq := &announcement_pb.DeleteAnnouncementRequest{}
	postData, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body format", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(postData, deleteReq)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body fields", http.StatusBadRequest)
		return
	}
	deleteReq.AccessToken = reqToken
	resp, err := s.announcement_client.DeleteAnnouncement(context.Background(), deleteReq)
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

func (s *Repository) EditPost(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Jwt ")
	if len(splitToken) < 2 {
		helpers.MessageGenerator(w, "User is UnAuthorized", http.StatusUnauthorized)
		return
	}
	reqToken = splitToken[1]
	editReq := &announcement_pb.EditPostRequest{}
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
	resp, err := s.announcement_client.EditPost(context.Background(), editReq)
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

func (s *Repository) UploadHostHouseImage(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Jwt ")
	if len(splitToken) < 2 {
		helpers.MessageGenerator(w, "User is UnAuthorized", http.StatusUnauthorized)
		return
	}
	reqToken = splitToken[1]
	// Parse the form data, including the uploaded file
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit for the image size
	if err != nil {
		helpers.MessageGenerator(w, "image size is too large", http.StatusBadRequest)
		return
	}
	uploadReq := &announcement_pb.HostHouseImageRequest{
		AccessToken: reqToken,
		ImageData:   make([][]byte, 0, 3),
	}
	for i := 1; i < 4; i++ {
		file, _, err := r.FormFile(fmt.Sprintf("image-%d", i))
		if err != nil {
			break
		}
		defer file.Close()
		image_data, err := io.ReadAll(file)
		if err != nil {
			helpers.MessageGenerator(w, "Image file is corrupted", http.StatusBadRequest)
			return
		}
		uploadReq.ImageData = append(uploadReq.ImageData, image_data)
	}
	resp, err := s.announcement_client.UploadHostHouseImage(context.Background(), uploadReq)
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
