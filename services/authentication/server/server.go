package server

import (
	"context"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/authentication/pb"
)

type Repository struct {
	pb.UnimplementedAuthServer
}

func New() *Repository {
	return &Repository{}
}

func (s *Repository) Signup(ctx context.Context, r *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	return nil, nil
}

func (s *Repository) Login(ctx context.Context, r *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, nil
}
