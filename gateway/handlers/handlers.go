package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/gateway/helpers"
	auth_pb "github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/gateway/services/authentication"
	"google.golang.org/grpc"
)

type Repository struct {
	auth_client auth_pb.AuthClient
}

func New(auth_conn *grpc.ClientConn) *Repository {
	return &Repository{
		auth_client: auth_pb.NewAuthClient(auth_conn),
	}
}

func (s *Repository) SignUp(w http.ResponseWriter, r *http.Request) {
	signupReq := &auth_pb.SignUpRequest{}
	postData, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body format", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(postData, signupReq)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body fields", http.StatusBadRequest)
		return
	}
	resp, err := s.auth_client.Signup(context.Background(), signupReq)
	if err != nil {
		helpers.MessageGenerator(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if resp.Message == "success" {
		helpers.MessageGenerator(w, resp.Message, http.StatusCreated)
	} else {
		helpers.MessageGenerator(w, resp.Message, http.StatusBadRequest)
	}
}

func (s *Repository) Login(w http.ResponseWriter, r *http.Request) {
	loginReq := &auth_pb.LoginRequest{}
	postData, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body format", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(postData, loginReq)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body fields", http.StatusBadRequest)
		return
	}
	resp, err := s.auth_client.Login(context.Background(), loginReq)
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

func (s *Repository) Refresh(w http.ResponseWriter, r *http.Request) {
	refreshReq := &auth_pb.RefreshRequest{}
	postData, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body format", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(postData, refreshReq)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body fields", http.StatusBadRequest)
		return
	}
	resp, err := s.auth_client.Refresh(context.Background(), refreshReq)
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

func (s *Repository) ForgetPassword(w http.ResponseWriter, r *http.Request) {
	ForgetPasswordReq := &auth_pb.ForgetPasswordRequest{}
	postData, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body format", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(postData, ForgetPasswordReq)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body fields", http.StatusBadRequest)
		return
	}
	resp, err := s.auth_client.ForgetPassword(context.Background(), ForgetPasswordReq)
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
func (s *Repository) ResetPassword(w http.ResponseWriter, r *http.Request) {
	resetPasswordReq := &auth_pb.ResetPasswordRequest{}
	access_token := r.URL.Query().Get("token")
	postData, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body format", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(postData, resetPasswordReq)
	if err != nil {
		helpers.MessageGenerator(w, "wrong post body fields", http.StatusBadRequest)
		return
	}
	resetPasswordReq.AccessToken = access_token
	resp, err := s.auth_client.ResetPassword(context.Background(), resetPasswordReq)
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
func (s *Repository) EmailVerification(w http.ResponseWriter, r *http.Request) {
	verifyReq := &auth_pb.VerifyRequest{}
	access_token := r.URL.Query().Get("token")
	verifyReq.Token = access_token
	resp, err := s.auth_client.EmailVerification(context.Background(), verifyReq)
	if err != nil {
		helpers.MessageGenerator(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if resp.Message == "user verified successfully" {
		helpers.MessageGenerator(w, resp.Message, http.StatusOK)
	} else {
		helpers.MessageGenerator(w, resp.Message, http.StatusBadRequest)
	}
}
