package db

import "github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/announcement/pb"

type Repository interface {
	GetIdFromUsername(username string) (int, error)
	GetUsernameFromId(id int) (string, error)
	InsertAnnouncement(announcementInfo *pb.CreateCardRequest, user_id int) (int, error)
	InsertAnnouncementLanguage(announcement_id int, lang string) error
	CheckAnnouncementTimeValidation(startDate string, endDate string, user_id int) (bool, error)
	GetAnnouncementDetails(filter []string, sort string, pagesize, pagenumber int) (*pb.GetCardResponse, error)
	GetLanguagesOfAnnouncement(announcement_id int) ([]string, error)
	InsertOffer(offerInfo *pb.CreateOfferRequest, user_id int) error
	GetGuestID(announcementID int) (int, error)
	GetOfferDetails(announcement_id int) (*pb.GetOfferResponse, error)
	GetProfileAnnouncementDetails(user_id int) (*pb.GetCardProfileResponse, error)
	ValidateOffer(announcement_id int, user_id int) (bool, string, error)
	InsertPost(postInfo *pb.CreatePostRequest) error
	GetLastPostId() (int, error)
	GetMyPostDetails(guest_id int) (*pb.GetMyPostResponse, error)
	GetPostHostDetails(username string) (*pb.GetPostHostResponse, error)
	AcceptUserAsHost(offerInfo *pb.AcceptOfferRequest) error
	RejectUserOffer(offerInfo *pb.RejectOfferRequest) error
	UpdateAnnouncementInformation(announcementInfo *pb.EditAnnouncementRequest) error
	DeleteAnnouncement(announcement_id int) error
	UpdatePostInformation(postInfo *pb.EditPostRequest) error
	GetHostId(announcement_id int) (int, error)
	UpdateHostImagesCount(user_id, imageCount int) error
	DeleteUserChatList(announcement_id, host_id int) error
	UpdateChatListStatus(announcement_id, host_id int) error
	UpdateAnnouncementStatus(announcement_id, host_id int) error
	UpdateMainHostStatusInChatList(announcement_id, host_id int) error
}
