package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/chat/models"
)

type PostgresRepository struct {
	DB *sql.DB
}

func NewPostgresRepository(db *sql.DB) Repository {
	return &PostgresRepository{
		DB: db,
	}
}

func (s *PostgresRepository) InsertChatMessage(message *models.Message) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := `insert into messages (message , user_id , username , time , chat_id) values($1,$2,$3,$4,$5)`
	_, err := s.DB.ExecContext(ctx, query, message.Message, message.UserID, message.UserName, message.Time, message.ChatID)
	if err != nil {
		log.Println(err.Error())
	}
}
