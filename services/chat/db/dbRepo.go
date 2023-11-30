package db

import "github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/chat/models"

type Repository interface {
	InsertChatMessage(message *models.Message)
	GetHistory(chat_id string, PageNumber, numberOfMessages int) ([]models.ChatMessage, error)
	GetCount(chat_id string) int
}
