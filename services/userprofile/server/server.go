package server

import (
	"context"
	"errors"
	"log"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/userprofile/db"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/userprofile/helper"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/userprofile/pb"
	"golang.org/x/crypto/bcrypt"
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

func (s *Repository) EditProfile(ctx context.Context, r *pb.EditProfileRequest) (*pb.EditProfileResponse, error) {
	claims, err := helper.DecodeToken(r.AccessToken)
	if err != nil {
		resp := &pb.EditProfileResponse{
			Message: "User is UnAuthorized",
		}
		return resp, nil
	}
	if r.CurrentPassword != "" {
		current_password, err := s.DB.GetUserPassword(claims.UserName)
		if err != nil {
			log.Println(err.Error())
			err := errors.New("internal error while getting user password from database - userprofile service")
			return nil, err
		}
		err = bcrypt.CompareHashAndPassword([]byte(current_password), []byte(r.CurrentPassword))
		if err != nil {
			resp := &pb.EditProfileResponse{
				Message: "Wrong current password",
			}
			return resp, nil
		}
	}
	err = s.DB.UpdateUserInformation(claims.UserName, r)
	if err != nil {
		log.Println(err.Error())
		err := errors.New("internal error while updating user info - userprofile service")
		return nil, err
	}
	resp := &pb.EditProfileResponse{
		Message: "success",
	}
	return resp, nil
}

func (s *Repository) UserNotifications(ctx context.Context, r *pb.UserNotificationRequest) (*pb.UserNotificationResponse, error) {
	claims, err := helper.DecodeToken(r.AccessToken)
	if err != nil {
		log.Println(err.Error())
		resp := &pb.UserNotificationResponse{
			Message: err.Error(),
		}
		return resp, nil
	}

	notfis, err := s.DB.GetNotificationByID(claims.UserID)
	if err != nil {
		log.Println(err.Error())
		err := errors.New("internal error while getting user notifications - userprofile service")
		return nil, err
	}

	if len(notfis) == 0 {
		resp := &pb.UserNotificationResponse{
			Message: "There are no notifications",
		}
		return resp, nil
	}

	resp := &pb.UserNotificationResponse{
		Notifs:  notfis,
		Message: "success",
	}
	return resp, nil
}
