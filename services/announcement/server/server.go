package server

import (
	"bytes"
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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/viper"
)

type Repository struct {
	pb.UnimplementedAnnouncementServer
	DB    db.Repository
	Queue *queue.Queue
	S3    *s3.S3
}

func New(db db.Repository, q *queue.Queue, s3 *s3.S3) *Repository {
	return &Repository{
		DB:    db,
		Queue: q,
		S3:    s3,
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
		username, err := s.DB.GetUsernameFromId(int(card.UserId))
		if err != nil {
			respErr := errors.New("internal server error while getting username from id - announcement service")
			log.Println(err)
			return nil, respErr
		}
		card.UserUsername = username
		card.Image, err = helper.GetImageURL(s.S3, fmt.Sprintf("user-%d", card.UserId))
		if err != nil {
			log.Println(err.Error())
			err := errors.New("internal error while getting user image from object storage - announcement service")
			return nil, err
		}
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

	check, message, err := s.DB.ValidateOffer(int(r.AnnouncementId), claims.UserID)
	if err != nil {
		respErr := errors.New("internal server error while validating new offer - announcement service")
		log.Println(err)
		return nil, respErr
	}
	if !check {
		resp := &pb.CreateOfferResponse{
			Message: message,
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

func (s *Repository) GetOffer(ctx context.Context, r *pb.GetOfferRequest) (*pb.GetOfferResponse, error) {
	_, err := helper.DecodeToken(r.AccessToken)
	if err != nil {
		resp := &pb.GetOfferResponse{
			Message: "User is UnAuthorized - announcement service",
		}
		return resp, nil
	}
	resp, err := s.DB.GetOfferDetails(int(r.AnnouncementId))
	if err != nil {
		log.Println(err.Error())
		err := errors.New("internal error while getting offer info - announcement service")
		return nil, err
	}
	for _, offer := range resp.Offers {
		offer.Image, err = helper.GetImageURL(s.S3, fmt.Sprintf("user-%d", offer.HostId))
		if err != nil {
			log.Println(err.Error())
			err := errors.New("internal error while getting user image from object storage - announcement service")
			return nil, err
		}
	}
	resp.Message = "success"
	return resp, nil
}

func (s *Repository) GetCardProfile(ctx context.Context, r *pb.GetCardProfileRequest) (*pb.GetCardProfileResponse, error) {
	claims, err := helper.DecodeToken(r.AccessToken)
	if err != nil {
		resp := &pb.GetCardProfileResponse{
			Message: "User is UnAuthorized - announcement service",
		}
		return resp, nil
	}
	resp, err := s.DB.GetProfileAnnouncementDetails(claims.UserID)
	if err != nil {
		log.Println(err.Error())
		err := errors.New("internal error while getting profile announcement info - announcement service")
		return nil, err
	}
	for _, card := range resp.Cards {
		languages, err := s.DB.GetLanguagesOfAnnouncement(int(card.CardId))
		if err != nil {
			respErr := errors.New("internal server error while adding new profile announcement language - announcement service")
			log.Println(err)
			return nil, respErr
		}
		card.PreferredLanguages = languages[:]
	}
	resp.Message = "success"
	return resp, nil
}

func (s *Repository) CreatePost(ctx context.Context, r *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	_, err := helper.DecodeToken(r.AccessToken)
	if err != nil {
		resp := &pb.CreatePostResponse{
			Message: "User is UnAuthorized - announcement service",
		}
		return resp, nil
	}

	err = s.DB.InsertPost(r)
	if err != nil {
		respErr := errors.New("internal server error while adding new post - announcement service")
		log.Println(err)
		return nil, respErr
	}

	resp := pb.CreatePostResponse{
		Message: "success",
	}
	return &resp, nil
}

func (s *Repository) UploadPostImage(ctx context.Context, r *pb.PostImageRequest) (*pb.PostImageResponse, error) {
	_, err := helper.DecodeToken(r.AccessToken)
	if err != nil {
		resp := &pb.PostImageResponse{
			Message: "User is UnAuthorized - announcement service",
		}
		return resp, nil
	}
	// Get last post id
	post_id, err := s.DB.GetLastPostId()
	if err != nil {
		respErr := errors.New("internal server error while getting last post id - announcement service")
		return nil, respErr
	}
	bucketName := viper.Get("OBJECT_STORAGE_BUCKET_NAME").(string)
	// Upload the image to S3
	_, err = s.S3.PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(bucketName),
		Key:                aws.String(fmt.Sprintf("post-%d", post_id)),
		ACL:                aws.String("private"), // Set ACL as needed
		Body:               bytes.NewReader(r.ImageData),
		ContentLength:      aws.Int64(int64(len(r.ImageData))),
		ContentType:        aws.String("image/jpeg"), // Set content type based on your file type
		ContentDisposition: aws.String("attachment"),
	})
	if err != nil {
		log.Println(err.Error())
		err := errors.New("internal error while uploading image to object storage - announcement service")
		return nil, err
	}
	resp := &pb.PostImageResponse{
		Message: "success",
	}
	return resp, nil
}

func (s *Repository) GetMyPost(ctx context.Context, r *pb.GetMyPostRequest) (*pb.GetMyPostResponse, error) {
	claims, err := helper.DecodeToken(r.AccessToken)
	if err != nil {
		resp := &pb.GetMyPostResponse{
			Message: "User is UnAuthorized - announcement service",
		}
		return resp, nil
	}

	resp, err := s.DB.GetMyPostDetails(claims.UserID)
	if err != nil {
		log.Println(err.Error())
		err := errors.New("internal error while getting post info - announcement service")
		return nil, err
	}

	for _, post := range resp.Posts {
		post.PostImage, err = helper.GetImageURL(s.S3, fmt.Sprintf("post-%d", post.PostId))
		if err != nil {
			log.Println(err.Error())
			err := errors.New("internal error while getting post image from object storage - announcement service")
			return nil, err
		}
		post.GuestImage, err = helper.GetImageURL(s.S3, fmt.Sprintf("user-%d", post.GuestId))
		if err != nil {
			log.Println(err.Error())
			err := errors.New("internal error while getting user image from object storage - announcement service")
			return nil, err
		}
		post.HostImage, err = helper.GetImageURL(s.S3, fmt.Sprintf("post-%d", post.HostId))
		if err != nil {
			log.Println(err.Error())
			err := errors.New("internal error while getting user image from object storage - announcement service")
			return nil, err
		}
	}

	resp.Message = "success"
	return resp, nil
}

func (s *Repository) GetPostHost(ctx context.Context, r *pb.GetPostHostRequest) (*pb.GetPostHostResponse, error) {
	_, err := helper.DecodeToken(r.AccessToken)
	if err != nil {
		resp := &pb.GetPostHostResponse{
			Message: "User is UnAuthorized - announcement service",
		}
		return resp, nil
	}

	resp, err := s.DB.GetPostHostDetails(int(r.HostId))
	if err != nil {
		log.Println(err.Error())
		err := errors.New("internal error while getting post info - announcement service")
		return nil, err
	}

	for _, post := range resp.Posts {
		post.PostImage, err = helper.GetImageURL(s.S3, fmt.Sprintf("post-%d", post.PostId))
		if err != nil {
			log.Println(err.Error())
			err := errors.New("internal error while getting post image from object storage - announcement service")
			return nil, err
		}
		post.GuestImage, err = helper.GetImageURL(s.S3, fmt.Sprintf("user-%d", post.GuestId))
		if err != nil {
			log.Println(err.Error())
			err := errors.New("internal error while getting user image from object storage - announcement service")
			return nil, err
		}
		post.HostImage, err = helper.GetImageURL(s.S3, fmt.Sprintf("post-%d", post.HostId))
		if err != nil {
			log.Println(err.Error())
			err := errors.New("internal error while getting user image from object storage - announcement service")
			return nil, err
		}
	}
	resp.Message = "success"
	return resp, nil
}

func (s *Repository) AcceptOffer(ctx context.Context, r *pb.AcceptOfferRequest) (*pb.AcceptOfferResponse, error) {
	_, err := helper.DecodeToken(r.AccessToken)
	if err != nil {
		resp := &pb.AcceptOfferResponse{
			Message: "User is UnAuthorized - announcement service",
		}
		return resp, nil
	}

	err = s.DB.AcceptUserAsHost(r)
	if err != nil {
		respErr := errors.New("internal server error while accepting offer - announcement service")
		log.Println(err)
		return nil, respErr
	}

	resp := pb.AcceptOfferResponse{
		Message: "success",
	}
	return &resp, nil
}

func (s *Repository) RejectOffer(ctx context.Context, r *pb.RejectOfferRequest) (*pb.RejectOfferResponse, error) {
	_, err := helper.DecodeToken(r.AccessToken)
	if err != nil {
		resp := &pb.RejectOfferResponse{
			Message: "User is UnAuthorized - announcement service",
		}
		return resp, nil
	}

	err = s.DB.RejectUserAsHost(r)
	if err != nil {
		respErr := errors.New("internal server error while rejecting offer - announcement service")
		log.Println(err)
		return nil, respErr
	}

	resp := pb.RejectOfferResponse{
		Message: "success",
	}
	return &resp, nil
}

func (s *Repository) EditAnnouncement(ctx context.Context, r *pb.EditAnnouncementRequest) (*pb.EditAnnouncementResponse, error) {
	claims, err := helper.DecodeToken(r.AccessToken)
	if err != nil {
		resp := &pb.EditAnnouncementResponse{
			Message: "User is UnAuthorized",
		}
		return resp, nil
	}
	check, err := s.DB.CheckAnnouncementTimeValidation(r.StartDate, r.EndDate, claims.UserID)
	if err != nil {
		respErr := errors.New("internal server error while checking announcement existance - announcement service")
		log.Println(err)
		return nil, respErr
	}
	if !check {
		resp := pb.EditAnnouncementResponse{
			Message: "you already have created an announcement in this time range",
		}
		return &resp, nil
	}
	err = s.DB.UpdateAnnouncementInformation(r)
	if err != nil {
		log.Println(err.Error())
		err := errors.New("internal error while updating announcement info - announcement service")
		return nil, err
	}
	for _, lang := range r.PreferredLanguages {
		err := s.DB.InsertAnnouncementLanguage(int(r.CardId), lang)
		if err != nil {
			respErr := errors.New("internal server error while adding new announcement language - announcement service")
			return nil, respErr
		}
	}
	resp := &pb.EditAnnouncementResponse{
		Message: "success",
	}
	return resp, nil
}