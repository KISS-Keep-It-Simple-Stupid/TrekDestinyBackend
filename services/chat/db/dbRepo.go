package db

import "github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/chat/models"

type Repository interface {
	InsertChatMessage(message *models.Message )
}
