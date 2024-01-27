package db

import (
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/userprofile/pb"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Repository interface {
	GetUserDetails(usename string) (*pb.ProfileDetailsResponse, error)
	UpdateUserInformation(username string, userInfo *pb.EditProfileRequest) error
	CheckUserExistance(userUserName string) (bool, error)
	GetUserPassword(username string) (string, error)
	GetPublicProfile(username string) (*pb.PublicProfileResponse, int, error)
	GetPublicProfileHost(guest_id int, host_username string) (*pb.PublicProfileHostResponse, string, error)
	GetIdFromUsername(username string) (int, error)
	InsertChatList(host_id, guest_id, announcement_id int) error
	GetChatList(guest_id int, obj *s3.S3) (*pb.ChatListResponse, error)
	GetHouseImagesCount(user_id int) (int, error)
	UpdateOfferStatus(announcement_id, host_id int) error
	InsertUserLanguage(user_id int, lang string) error
	InsertUserInterest(user_id int, interest string) error
	GetLanguagesOfUser(user_id int) ([]string, error)
	GetInterestsOfUser(user_id int) ([]string, error)
}
