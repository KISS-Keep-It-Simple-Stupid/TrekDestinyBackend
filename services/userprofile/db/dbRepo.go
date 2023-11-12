package db

import "github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/userprofile/pb"

type Repository interface {
	GetUserDetails(usename string) (*pb.ProfileDetailsResponse, error)
	UpdateUserInformation(username string, userInfo *pb.EditProfileRequest) error
	CheckUserExistance(userUserName string) (bool, error)
	GetUserPassword(username string) (string, error) 
}
