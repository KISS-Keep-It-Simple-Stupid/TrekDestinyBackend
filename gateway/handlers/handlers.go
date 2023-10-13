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
		helpers.RespGenerator(w, "wrong post body format", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(postData, signupReq)
	if err != nil {
		helpers.RespGenerator(w, "wrong post body fields", http.StatusBadRequest)
		return
	}
	resp, err := s.auth_client.Signup(context.Background(), signupReq)
	if err != nil {
		helpers.RespGenerator(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if resp.Message == "success" {
		helpers.RespGenerator(w, resp.Message, http.StatusCreated)
	} else {
		helpers.RespGenerator(w, resp.Message, http.StatusBadRequest)
	}
}
