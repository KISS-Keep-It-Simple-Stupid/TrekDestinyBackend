package db

import "github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/userprofile/pb"

type Repository interface {
	GetUserDetails(usename string) (*pb.ProfileDetailsResponse, error)
	UpdateUserInformation(username string, userInfo *pb.EditProfileRequest) error
	CheckUserExistance(userUserName string) (bool, error)
	GetUserPassword(username string) (string, error)
	GetPublicProfile(username string) (*pb.PublicProfileResponse, int, error)
	GetPublicProfileHost(guest_id int, host_username string) (*pb.PublicProfileHostResponse, string, error)
	GetIdFromUsername(username string) (int, error)
}
