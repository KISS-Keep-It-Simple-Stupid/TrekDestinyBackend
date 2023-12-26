package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/notification/models"
)

type postgresRepo struct {
	db *sql.DB
}

func New(dbParam *sql.DB) Repository {
	return &postgresRepo{
		db: dbParam,
	}
}

func (repo *postgresRepo) InsertNotification(message models.NotifMessage) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := `insert into notifications (user_id,message) 
			 values($1,$2)`
	_, err := repo.db.ExecContext(ctx, query,
		message.UserID,
		message.Message,
	)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return nil
}

func (s *postgresRepo) GetNotificationByID(userid int) ([]*models.NotifResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := `select id , message from notifications where user_id = $1`
	rows, err := s.db.QueryContext(ctx, query, userid)
	if err != nil {
		if err == sql.ErrNoRows {
			return make([]*models.NotifResponse, 0), nil
		}
		return nil, err
	}

	notifs := make([]*models.NotifResponse, 0)

	for rows.Next() {
		temp_notif := models.NotifResponse{}
		err := rows.Scan(&temp_notif.ID, &temp_notif.Message)
		if err != nil {
			return nil, err
		}
		notifs = append(notifs, &temp_notif)
	}

	return notifs, nil
}

func (s *postgresRepo) DeleteNotification(notfiID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := `delete from notifications where id = $1`
	_, err := s.db.ExecContext(ctx, query, notfiID)
	if err != nil {
		log.Println(err)
	}
	return err
}


