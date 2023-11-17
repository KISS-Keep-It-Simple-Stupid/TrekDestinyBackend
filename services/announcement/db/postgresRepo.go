package db

import (
	"context"
	"database/sql"
	"log"
	"math"
	"strings"
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

func (s *PostgresRepository) GetUsernameFromId(id int) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	var username string
	query := `select username from members where id = $1`
	err := s.DB.QueryRowContext(ctx, query, id).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}

func (s *PostgresRepository) InsertAnnouncement(announcementInfo *pb.CreateCardRequest, user_id int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := `insert into announcement (user_id, description, startdate, enddate, city, state, country, numberoftravelers) values ($1, $2, $3, $4, $5, $6, $7, $8)`
	startdate, err := time.Parse("2006-01-02", announcementInfo.StartDate)
	if err != nil {
		return -1, err
	}
	enddate, err := time.Parse("2006-01-02", announcementInfo.EndDate)
	if err != nil {
		return -1, err
	}
	_, err = s.DB.ExecContext(ctx, query, user_id, announcementInfo.Description, startdate, enddate, announcementInfo.DestinationCity, announcementInfo.DestinationState, announcementInfo.DestinationCountry, int(announcementInfo.NumberOfTravelers))
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

func (s *PostgresRepository) GetAnnouncementDetails(filter []string, sort string, pagesize, pagenumber int) (*pb.GetCardResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	resp := pb.GetCardResponse{}
	query := "select id, user_id, description, startdate, enddate, city, state, country, numberoftravelers from announcement"
	if !(len(filter) == 1 && filter[0] == "") {
		for i, singlefilter := range filter {
			parts := strings.Split(singlefilter, ":")
			field, value := parts[0], parts[1]
			if i != 0 {
				query += " and "
			} else {
				query += " where "
			}
			query += field + " = '" + value + "'"
		}
	}
	if sort != "" {
		parts := strings.Split(sort, ".")
		sortvalue, order := parts[0], parts[1]
		query += " order by " + sortvalue
		if order == "desc" {
			query += " desc"
		} else if order == "asc" {
			query += " asc"
		}
	}
	query += " limit $1 offset $2"
	rows, err := s.DB.QueryContext(ctx, query, pagesize, (pagenumber-1)*pagesize)
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
			&card.DestinationState,
			&card.DestinationCountry,
			&card.NumberOfTravelers)
		if err != nil {
			return nil, err
		}
		card.StartDate = startdate.Format("2006-01-02")
		card.EndDate = enddate.Format("2006-01-02")
		resp.Cards = append(resp.Cards, &card)
	}
	var cardcount int
	query = `select COUNT(*) from announcement`
	err = s.DB.QueryRow(query).Scan(&cardcount)
	if err != nil {
		return nil, err
	}
	resp.PageCount = int32(math.Ceil(float64(cardcount) / float64(pagesize)))
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

func (s *PostgresRepository) InsertOffer(offerInfo *pb.CreateOfferRequest, user_id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := `insert into announcement_offer (announcement_id, host_id) values ($1, $2)`
	_, err := s.DB.ExecContext(ctx, query, int(offerInfo.AnnouncementId), user_id)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (s *PostgresRepository) GetGuestID(announcementID int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := "select user_id from announcement where id = $1"
	guestID := 0
	err := s.DB.QueryRowContext(ctx, query, announcementID).Scan(&guestID)
	if err != nil {
		return -1, err
	}
	return guestID, nil
}

func (s *PostgresRepository) GetOfferDetails(announcement_id int) (*pb.GetOfferResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	resp := pb.GetOfferResponse{}
	query := "select host_id from announcement_offer where announcement_id = $1"
	rows, err := s.DB.QueryContext(ctx, query, announcement_id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		offer := pb.OfferRecord{}
		var firstname string
		var lastname string
		var username string
		err := rows.Scan(
			&offer.HostId)
		if err != nil {
			return nil, err
		}
		query := "select firstname, lastname, username from members where id = $1"
		err = s.DB.QueryRowContext(ctx, query, offer.HostId).Scan(&firstname, &lastname, &username)
		if err != nil {
			return nil, err
		}
		offer.HostFirstName = firstname
		offer.HostLastName = lastname
		offer.HostUsername = username
		resp.Offers = append(resp.Offers, &offer)
	}
	return &resp, nil
}

func (s *PostgresRepository) GetProfileAnnouncementDetails(user_id int) (*pb.GetCardProfileResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	resp := pb.GetCardProfileResponse{}
	query := "select id, user_id, description, startdate, enddate, city, state, country, numberoftravelers from announcement where user_id = $1"
	rows, err := s.DB.QueryContext(ctx, query, user_id)
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
			&card.DestinationState,
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

func (s *PostgresRepository) ValidateOffer(announcement_id int, user_id int) (bool, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	var temp_id int
	query := `select user_id from announcement where id = $1`
	err := s.DB.QueryRowContext(ctx, query, announcement_id).Scan(&temp_id)
	if err != nil {
		return true, "", err
	}
	if temp_id == user_id {
		return false, "you can not offer to your own announcement", nil
	}

	query = `select host_id from announcement_offer where announcement_id = $1`
	rows, err := s.DB.QueryContext(ctx, query, announcement_id)
	if err != nil {
		return true, "", err
	}
	for rows.Next() {
		err := rows.Scan(&temp_id)
		if err != nil {
			return true, "", err
		}
		if temp_id == user_id {
			return false, "you have already offered to this announcement", nil
		}
	}
	return true, "", nil
}
