package handlers

import (
	announcement_pb "github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/gateway/services/announcement"
	auth_pb "github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/gateway/services/authentication"
	userprofile_pb "github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/gateway/services/userprofile"
	"google.golang.org/grpc"
)

type Repository struct {
	auth_client         auth_pb.AuthClient
	userprofile_client  userprofile_pb.UserProfileClient
	announcement_client announcement_pb.AnnouncementClient
}

func New(auth_conn, userprofile_conn, announcement_conn *grpc.ClientConn) *Repository {
	return &Repository{
		auth_client:         auth_pb.NewAuthClient(auth_conn),
		userprofile_client:  userprofile_pb.NewUserProfileClient(userprofile_conn),
		announcement_client: announcement_pb.NewAnnouncementClient(announcement_conn),
	}
}
