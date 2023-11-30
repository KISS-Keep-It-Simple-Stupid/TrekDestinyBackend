package db

import (
	"context"
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/chat/models"
	"github.com/spf13/viper"
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

func (s *PostgresRepository) GetCount(chat_id string) int {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select count(*) from messages where chat_id = $1`
	var count int
	_ = s.DB.QueryRowContext(ctx, query, chat_id).Scan(&count)
	return count

}

func (s *PostgresRepository) GetHistory(chat_id string, pageNumber, numberOfMessages int) ([]models.ChatMessage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	pageSize, _ := strconv.Atoi(viper.Get("PAGESIZE").(string))
	query := `SELECT  user_id , username , message , time
				FROM    ( SELECT    ROW_NUMBER() OVER ( ORDER BY chat_id ) AS RowNum, *
						FROM      messages
						WHERE     chat_id = $1
						) AS RowConstrainedResult
				WHERE   RowNum > $2
					AND RowNum <= $3
				ORDER BY RowNum DESC`

	start := numberOfMessages - pageSize*pageNumber
	end := numberOfMessages - pageSize*(pageNumber-1)
	if start < 0 {
		start = 0
	}
	if end > 0 {
		rows, err := s.DB.QueryContext(ctx, query, chat_id, start, end)
		result := make([]models.ChatMessage, 0)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			temp := models.ChatMessage{}
			rows.Scan(&temp.UserID, &temp.UserName, &temp.Message, &temp.Time)
			result = append(result, temp)
		}
		return result, nil
	}
	return []models.ChatMessage{}, nil
}
