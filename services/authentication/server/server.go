package server

import (
	"context"
	"errors"
	"log"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/authentication/db"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/authentication/pb"
)

type Repository struct {
	pb.UnimplementedAuthServer
	DB db.Repository
}

func New(db db.Repository) *Repository {
	return &Repository{
		DB: db,
	}
}

func (s *Repository) Signup(ctx context.Context, r *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	check, err := s.DB.CheckUserExistance(r.Email, r.UserName)
	if err != nil {
		respErr := errors.New("internal server error while checking user existance - authentication service")
		log.Println(err)
		return nil, respErr
	}
	if check {
		resp := pb.SignUpResponse{
			Message: "user already exists",
		}
		return &resp, nil
	}
	err = s.DB.InsertUser(r)
	if err != nil {
		respErr := errors.New("internal server error while adding new user - authentication service")
		log.Println(err)
		return nil, respErr
	}
	resp := pb.SignUpResponse{
		Message: "success",
	}
	return &resp, nil
}

func (s *Repository) Login(ctx context.Context, r *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, nil
}
