package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
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
