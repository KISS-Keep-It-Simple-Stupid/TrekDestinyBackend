package server

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/userprofile/db"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/userprofile/helper"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/userprofile/pb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type Repository struct {
	pb.UnimplementedUserProfileServer
	DB db.Repository
	S3 *s3.S3
}

func New(db db.Repository, s3 *s3.S3) *Repository {
	return &Repository{
		DB: db,
		S3: s3,
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
	resp.Image, err = helper.GetImageURL(s.S3, fmt.Sprintf("user-%d", claims.UserID))
	if err != nil {
		log.Println(err.Error())
		err := errors.New("internal error while getting user image from object storage - userprofile service")
		return nil, err
	}

	resp.HostHouseImages = make([]string, 0, 3);
	for i := 1; i < 4; i++ {
		url, err := helper.GetImageURL(s.S3, fmt.Sprintf("user-%d-host-%d", claims.UserID, i))
		if err != nil {
			break
		}
		resp.HostHouseImages = append(resp.HostHouseImages, url)
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

func (s *Repository) UploadImage(ctx context.Context, r *pb.ImageRequest) (*pb.ImageResponse, error) {
	claims, err := helper.DecodeToken(r.AccessToken)
	if err != nil {
		resp := &pb.ImageResponse{
			Message: "User is UnAuthorized",
		}
		return resp, nil
	}
	bucketName := viper.Get("OBJECT_STORAGE_BUCKET_NAME").(string)
	// Upload the image to S3
	_, err = s.S3.PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(bucketName),
		Key:                aws.String(fmt.Sprintf("user-%d", claims.UserID)),
		ACL:                aws.String("private"), // Set ACL as needed
		Body:               bytes.NewReader(r.ImageData),
		ContentLength:      aws.Int64(int64(len(r.ImageData))),
		ContentType:        aws.String("image/jpeg"), // Set content type based on your file type
		ContentDisposition: aws.String("attachment"),
	})
	if err != nil {
		log.Println(err.Error())
		err := errors.New("internal error while uploading image to object storage - userprofile service")
		return nil, err
	}
	resp := &pb.ImageResponse{
		Message: "success",
	}
	return resp, nil
}

func (s *Repository) PublicProfile(ctx context.Context, r *pb.PublicProfileRequest) (*pb.PublicProfileResponse, error) {
	_, err := helper.DecodeToken(r.AccessToken)
	if err != nil {
		resp := &pb.PublicProfileResponse{
			Message: "User is UnAuthorized",
		}
		return resp, nil
	}

	resp, id, err := s.DB.GetPublicProfile(r.Username)
	if err != nil {
		log.Println(err.Error())
		err := errors.New("internal error while getting public profile details - userprofile service")
		return nil, err
	}
	resp.Message = "success"
	resp.Image, err = helper.GetImageURL(s.S3, fmt.Sprintf("user-%d", id))
	if err != nil {
		log.Println(err.Error())
		err := errors.New("internal error while getting image url - userprofile service")
		return nil, err
	}
	return resp, nil

}

func (s *Repository) PublicProfileHost(ctx context.Context, r *pb.PublicProfileHostRequest) (*pb.PublicProfileHostResponse, error) {
	claims, err := helper.DecodeToken(r.AccessToken)
	if err != nil {
		resp := &pb.PublicProfileHostResponse{
			Message: "User is UnAuthorized",
		}
		return resp, nil
	}
	resp, message, err := s.DB.GetPublicProfileHost(claims.UserID, r.Username)
	if err != nil {
		log.Println(err.Error())
		err := errors.New("internal error while getting public profile host details - userprofile service")
		return nil, err
	}
	if message != "" {
		resp.Message = message
		return resp, nil
	}
	user_id, err := s.DB.GetIdFromUsername(resp.UserName)
	if err != nil {
		return nil, err
	}
	resp.Image, err = helper.GetImageURL(s.S3, fmt.Sprintf("user-%d", user_id))
	if err != nil {
		log.Println(err.Error())
		err := errors.New("internal error while getting user image from object storage - userprofile service")
		return nil, err
	}

	host_id, err := s.DB.GetIdFromUsername(r.Username)
	if err != nil {
		return nil, err
	}

	resp.HostHouseImages = make([]string, 0, 3);
	for i := 1; i < 4; i++ {
		url, err := helper.GetImageURL(s.S3, fmt.Sprintf("user-%d-host-%d", host_id, i))
		if err != nil {
			break
		}
		resp.HostHouseImages = append(resp.HostHouseImages, url)
	}

	resp.Message = "success"
	return resp, nil
}

func (s *Repository) AddToChatList(ctx context.Context, r *pb.AddChatListRequest) (*pb.AddChatListResponse, error) {
	claims, err := helper.DecodeToken(r.AccessToken)
	if err != nil {
		resp := &pb.AddChatListResponse{
			Message: "User is UnAuthorized",
		}
		return resp, nil
	}
	err = s.DB.InsertChatList(int(r.HostID), claims.UserID, int(r.AnnouncementID))
	if err != nil {
		log.Println(err.Error())
		err := errors.New("internal error while inserting an item to chat list - userprofile service")
		return nil, err
	}
	resp := &pb.AddChatListResponse{
		Message: "success",
	}
	return resp, nil
}

func (s *Repository) GetChatList(ctx context.Context, r *pb.ChatListRequest) (*pb.ChatListResponse, error) {
	claims, err := helper.DecodeToken(r.AccessToken)
	if err != nil {
		resp := &pb.ChatListResponse{
			Message: "User is UnAuthorized",
		}
		return resp, nil
	}
	resp, err := s.DB.GetChatList(claims.UserID, s.S3)
	if err != nil {
		log.Println(err.Error())
		err := errors.New("internal error while retrieving chat list items - userprofile service")
		return nil, err
	}
	resp.Message = "success"
	return resp, nil
}
