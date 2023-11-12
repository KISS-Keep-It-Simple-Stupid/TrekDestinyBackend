package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/announcement/db"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/announcement/helper"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/announcement/models"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/announcement/pb"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/announcement/queue"
)

type Repository struct {
	pb.UnimplementedAnnouncementServer
	DB    db.Repository
	Queue *queue.Queue
}

func New(db db.Repository, q *queue.Queue) *Repository {
	return &Repository{
		DB:    db,
		Queue: q,
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

func (s *Repository) GetCard(ctx context.Context, r *pb.GetCardRequest) (*pb.GetCardResponse, error) {
	_, err := helper.DecodeToken(r.AccessToken)
	if err != nil {
		resp := &pb.GetCardResponse{
			Message: "User is UnAuthorized - announcement service",
		}
		return resp, nil
	}
	filter := strings.Split(r.FilterValues, ",")
	sort := r.SortValue
	pagesize, _ := strconv.Atoi(r.PageSize)
	pagenumber, _ := strconv.Atoi(r.PageNumber)
	resp, err := s.DB.GetAnnouncementDetails(filter, sort, pagesize, pagenumber)
	if err != nil {
		log.Println(err.Error())
		err := errors.New("internal error while getting announcement info - announcement service")
		return nil, err
	}
	for _, card := range resp.Cards {
		languages, err := s.DB.GetLanguagesOfAnnouncement(int(card.CardId))
		if err != nil {
			respErr := errors.New("internal server error while adding new announcement language - announcement service")
			log.Println(err)
			return nil, respErr
		}
		card.PreferredLanguages = languages[:]
	}
	resp.Message = "success"
	return resp, nil
}

func (s *Repository) CreateOffer(ctx context.Context, r *pb.CreateOfferRequest) (*pb.CreateOfferResponse, error) {
	claims, err := helper.DecodeToken(r.AccessToken)
	if err != nil {
		resp := &pb.CreateOfferResponse{
			Message: "User is UnAuthorized - announcement service",
		}
		return resp, nil
	}

	err = s.DB.InsertOffer(r, claims.UserID)
	if err != nil {
		respErr := errors.New("internal server error while adding new offer - announcement service")
		log.Println(err)
		return nil, respErr
	}

	guestID, err := s.DB.GetGuestID(int(r.AnnouncementId))
	if err != nil {
		respErr := errors.New("internal server error while getting guest id - announcement service")
		log.Println(err)
		return nil, respErr
	}

	go func(host string, guest int, q *queue.Queue) {
		notifMessage := models.NotificationMessage{
			UserID:  guest,
			Message: fmt.Sprintf("%s sent you a hosting request", host),
		}
		q.Send(&notifMessage)
	}(claims.UserName, guestID, s.Queue)
	
	resp := pb.CreateOfferResponse{
		Message: "success",
	}
	return &resp, nil
}
