package server

import (
	"context"
	"errors"
	"log"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/announcement/db"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/announcement/helper"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/announcement/pb"
)

type Repository struct {
	pb.UnimplementedAnnouncementServer
	DB db.Repository
}

func New(db db.Repository) *Repository {
	return &Repository{
		DB: db,
	}
}

func (s *Repository) CreateCard(ctx context.Context, r *pb.CreateCardRequest) (*pb.CreateCardResponse, error) {
	claims, err := helper.DecodeToken(r.AccessToken)
	if err != nil {
		resp := &pb.CreateCardResponse{
			Message: "User is UnAuthorized - announcement service",
		}
		return resp, nil
	}

	user_id, err := s.DB.GetIdFromUsername(claims.UserName)
	if err != nil {
		respErr := errors.New("internal server error while converting username to id - announcement service")
		log.Println(err)
		return nil, respErr
	}

	check, err := s.DB.CheckAnnouncementTimeValidation(r.StartDate, r.EndDate, user_id)
	if err != nil {
		respErr := errors.New("internal server error while checking announcement existance - announcement service")
		log.Println(err)
		return nil, respErr
	}
	if !check {
		resp := pb.CreateCardResponse{
			Message: "you already have created an announcement in this time range",
		}
		return &resp, nil
	}

	announcement_id, err := s.DB.InsertAnnouncement(r, user_id)
	if err != nil {
		respErr := errors.New("internal server error while adding new announcement - announcement service")
		log.Println(err)
		return nil, respErr
	}
	for _, lang := range r.PreferredLanguages {
		err := s.DB.InsertAnnouncementLanguage(announcement_id, lang)
		if err != nil {
			respErr := errors.New("internal server error while adding new announcement language - announcement service")
			log.Println(err)
			return nil, respErr
		}
	}
	resp := pb.CreateCardResponse{
		Message: "success",
	}
	return &resp, nil
}
