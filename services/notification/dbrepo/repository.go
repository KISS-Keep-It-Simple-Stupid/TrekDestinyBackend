package dbrepo

import "github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/notification/models"

type Repository interface {
	InsertNotification(message models.NotifMessage) error
	GetNotificationByID(userid int) ([]*models.NotifResponse, error)
	DeleteNotification(notifID int) error
}
