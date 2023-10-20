package server

import (
	"context"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/userprofile/db"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/userprofile/pb"
)

type Repository struct {
	pb.UnimplementedUserProfileServer
	DB db.Repository
}

func New(db db.Repository) *Repository {
	return &Repository{
		DB: db,
	}
}

func (s *Repository) ProfileDetails(ctx context.Context, r  *pb.ProfileDetailsRequest) (*pb.ProfileDetailsResponse, error){
	return nil , nil
}
