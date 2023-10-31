package db

import "github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/announcement/pb"

type Repository interface {
	GetIdFromUsername(username string) (int, error)
	InsertAnnouncement(announcementInfo *pb.CreateCardRequest, user_id int) (int, error)
	InsertAnnouncementLanguage(announcement_id int, lang string) error
	CheckAnnouncementTimeValidation(startDate string, endDate string, user_id int) (bool, error)
	GetAnnouncementDetails() (*pb.GetCardResponse, error)
	GetLanguagesOfAnnouncement(announcement_id int) ([]string, error)
}
