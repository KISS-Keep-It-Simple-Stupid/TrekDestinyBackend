package server

import (
	"context"
	"errors"
	"log"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/userprofile/db"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/userprofile/helper"
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

func (s *Repository) ProfileDetails(ctx context.Context, r *pb.ProfileDetailsRequest) (*pb.ProfileDetailsResponse, error) {
	claims, err := helper.DecodeToken(r.AccessToken)
	if err != nil {
		resp := &pb.ProfileDetailsResponse{
			Message: "User is UnAuthorized",
		}
		return resp, nil
	}
	resp, err := s.DB.GetUserDetails(claims.UserName)
	if err != nil {
		log.Println(err.Error())
		err := errors.New("internal error while getting user info - userprofile service")
		return nil, err
	}
	resp.Message = "success"
	return resp, nil
}
