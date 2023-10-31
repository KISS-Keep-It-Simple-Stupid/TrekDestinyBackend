package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/announcement/pb"
)

type PostgresRepository struct {
	DB *sql.DB
}

func NewPostgresRepository(db *sql.DB) Repository {
	return &PostgresRepository{
		DB: db,
	}
}

func (s *PostgresRepository) GetIdFromUsername(username string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	var id int
	query := `select id from members where username = $1`
	err := s.DB.QueryRowContext(ctx, query, username).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (s *PostgresRepository) InsertAnnouncement(announcementInfo *pb.CreateCardRequest, user_id int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := `insert into announcement (user_id, description, startdate, enddate, city, country, numberoftravelers) values ($1, $2, $3, $4, $5, $6, $7)`
	startdate, err := time.Parse("2006-01-02", announcementInfo.StartDate)
	if err != nil {
		return -1, err
	}
	enddate, err := time.Parse("2006-01-02", announcementInfo.EndDate)
	if err != nil {
		return -1, err
	}
	_, err = s.DB.ExecContext(ctx, query, user_id, announcementInfo.Description, startdate, enddate, announcementInfo.DestinationCity, announcementInfo.DestinationCountry, int(announcementInfo.NumberOfTravelers))
	if err != nil {
		return -1, err
	}
	var announcement_id int
    err = s.DB.QueryRow("select id from announcement order by id desc limit 1").Scan(&announcement_id)
    if err != nil {
        return -1, err
    }
	return int(announcement_id), nil
}

func (s *PostgresRepository) InsertAnnouncementLanguage(announcement_id int, lang string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := `insert into announcement_language (announcement_id, language) values ($1, $2)`
	_, err := s.DB.ExecContext(ctx, query, announcement_id, lang)
	return err
}

func (s *PostgresRepository) CheckAnnouncementTimeValidation(startDate string, endDate string, user_id int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := `select id from announcement where user_id = $1 and enddate >= $2 and startdate <= $3`
	startdate, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return false, err
	}
	enddate, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return false, err
	}
	var temp int
	err = s.DB.QueryRowContext(ctx, query, user_id, startdate, enddate).Scan(&temp)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

func (s *PostgresRepository) GetAnnouncementDetails() (*pb.GetCardResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	resp := pb.GetCardResponse{}
	query := "select id, user_id, description, startdate, enddate, city, country, numberoftravelers from announcement"
	rows, err := s.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		card := pb.CardRecord{}
		var startdate time.Time
		var enddate time.Time
		err := rows.Scan(
			&card.CardId, 
			&card.UserId, 
			&card.Description, 
			&startdate,
			&enddate,
			&card.DestinationCity,
			&card.DestinationCountry,
			&card.NumberOfTravelers)
		if err != nil {
			return nil, err
		}
		card.StartDate = startdate.Format("2006-01-02")
		card.EndDate = enddate.Format("2006-01-02")
		resp.Cards = append(resp.Cards, &card)
	}
	return &resp, nil
}

func (s *PostgresRepository) GetLanguagesOfAnnouncement(announcement_id int) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	var languages []string
	query := "select language from announcement_language where announcement_id = $1"
	rows, err := s.DB.QueryContext(ctx, query, announcement_id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var lang string
		err := rows.Scan(&lang)
		if err != nil {
			return nil, err
		}
		languages = append(languages, lang)
	}
	return languages, nil
}