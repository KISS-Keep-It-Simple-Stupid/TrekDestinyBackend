package db

import (
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/authentication/models"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/authentication/pb"
)

type Repository interface {
	InsertUser(user *pb.SignUpRequest) error
	CheckUserExistance(userEmail, userUserName string) (bool, error)
	GetLoginCridentials(userEmail string) (*models.LoginCridentials, bool, error)
}
