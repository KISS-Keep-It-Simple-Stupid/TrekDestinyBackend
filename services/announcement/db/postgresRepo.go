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
	query := "select id, user_id, description, startdate, enddate, city, state, country, numberoftravelers from announcement where announcement_status = $1"
	if !(len(filter) == 1 && filter[0] == "") {
		for _, singlefilter := range filter {
			parts := strings.Split(singlefilter, ":")
			field, value := parts[0], parts[1]
			query += " and "
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
	query2 := query + " limit $2 offset $3"
	rows, err := s.DB.QueryContext(ctx, query2, 1, pagesize, (pagenumber-1)*pagesize)
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
	rows, err = s.DB.QueryContext(ctx, query ,1 )
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		cardcount++
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
	query := "select host_id , offer_status from announcement_offer where announcement_id = $1"
	rows, err := s.DB.QueryContext(ctx, query, announcement_id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		offer := pb.OfferRecord{}
		var firstname string
		var lastname string
		var username string
		var offer_status int
		err := rows.Scan(&offer.HostId, &offer_status)
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
		offer.Status = int32(offer_status)
		resp.Offers = append(resp.Offers, &offer)
	}
	return &resp, nil
}

func (s *PostgresRepository) GetProfileAnnouncementDetails(user_id int) (*pb.GetCardProfileResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	resp := pb.GetCardProfileResponse{}
	query := "select id, user_id, description, startdate, enddate, city, state, country, numberoftravelers , announcement_status , main_host from announcement where user_id = $1"
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
			&card.NumberOfTravelers,
			&card.AnnouncementStatus,
			&card.MainHost)
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

func (s *PostgresRepository) InsertPost(postInfo *pb.CreatePostRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	var host_id int
	var guest_id int
	query := `select user_id, main_host from announcement where id = $1`
	err := s.DB.QueryRowContext(ctx, query, postInfo.AnnouncementId).Scan(&guest_id, &host_id)
	if err != nil {
		return err
	}
	query = `insert into post (announcement_id, host_id, guest_id, title, rating, body) values ($1, $2, $3, $4, $5, $6)`
	_, err = s.DB.ExecContext(ctx, query, postInfo.AnnouncementId, host_id, guest_id, postInfo.PostTitle, int(postInfo.HostRating), postInfo.PostBody)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresRepository) GetLastPostId() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	var post_id int
	err := s.DB.QueryRowContext(ctx, "select id from post order by id desc limit 1").Scan(&post_id)
	if err != nil {
		return -1, err
	}
	return post_id, nil
}

func (s *PostgresRepository) GetMyPostDetails(guest_id int) (*pb.GetMyPostResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	resp := pb.GetMyPostResponse{}
	query := "select id, announcement_id, host_id, guest_id, title, rating, body from post where guest_id = $1"
	rows, err := s.DB.QueryContext(ctx, query, guest_id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		post := pb.PostRecord{}
		var host_id int
		var guest_id int
		err := rows.Scan(
			&post.PostId,
			&post.AnnouncementId,
			&host_id,
			&guest_id,
			&post.PostTitle,
			&post.HostRating,
			&post.PostBody)
		if err != nil {
			return nil, err
		}
		var host_username string
		var guest_username string
		query := "select username from members where id = $1"
		err = s.DB.QueryRowContext(ctx, query, host_id).Scan(&host_username)
		if err != nil {
			return nil, err
		}
		err = s.DB.QueryRowContext(ctx, query, guest_id).Scan(&guest_username)
		if err != nil {
			return nil, err
		}
		post.HostUsername = host_username
		post.GuestUsername = guest_username
		post.HostId = int32(host_id)
		post.GuestId = int32(guest_id)
		resp.Posts = append(resp.Posts, &post)
	}
	return &resp, nil
}

func (s *PostgresRepository) GetPostHostDetails(host_id int) (*pb.GetPostHostResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	resp := pb.GetPostHostResponse{}
	query := "select id, announcement_id, host_id, guest_id, title, rating, body from post where host_id = $1"
	rows, err := s.DB.QueryContext(ctx, query, host_id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		post := pb.PostRecord{}
		var host_id int
		var guest_id int
		err := rows.Scan(
			&post.PostId,
			&post.AnnouncementId,
			&host_id,
			&guest_id,
			&post.PostTitle,
			&post.HostRating,
			&post.PostBody)
		if err != nil {
			return nil, err
		}
		var host_username string
		var guest_username string
		query := "select username from members where id = $1"
		err = s.DB.QueryRowContext(ctx, query, host_id).Scan(&host_username)
		if err != nil {
			return nil, err
		}
		err = s.DB.QueryRowContext(ctx, query, guest_id).Scan(&guest_username)
		if err != nil {
			return nil, err
		}
		post.HostUsername = host_username
		post.GuestUsername = guest_username
		post.HostId = int32(host_id)
		post.GuestId = int32(guest_id)
		resp.Posts = append(resp.Posts, &post)
	}
	return &resp, nil
}

func (s *PostgresRepository) AcceptUserAsHost(offerInfo *pb.AcceptOfferRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := `update announcement set main_host = $1, status = 'Accepted' where id = $2`
	_, err := s.DB.ExecContext(ctx, query, offerInfo.HostId, offerInfo.AnnouncementId)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresRepository) RejectUserOffer(offerInfo *pb.RejectOfferRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := `delete from announcement_offer where announcement_id = $1 and host_id = $2`
	_, err := s.DB.ExecContext(ctx, query, offerInfo.AnnouncementId, offerInfo.HostId)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresRepository) UpdateAnnouncementInformation(announcementInfo *pb.EditAnnouncementRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	startdate, err := time.Parse("2006-01-02", announcementInfo.StartDate)
	if err != nil {
		return err
	}
	enddate, err := time.Parse("2006-01-02", announcementInfo.EndDate)
	if err != nil {
		return err
	}
	query := `update announcement set  
		description = COALESCE(NULLIF($1, ''), description),
		startdate = $2,
		enddate = $3,
		city = COALESCE(NULLIF($4, ''), city),
		state = COALESCE(NULLIF($5, ''), state),
		country = COALESCE(NULLIF($6, ''), country),
		numberoftravelers = COALESCE(NULLIF($7, 0), numberoftravelers)
		where id = $8`
	_, err = s.DB.ExecContext(ctx, query,
		announcementInfo.Description,
		startdate,
		enddate,
		announcementInfo.DestinationCity,
		announcementInfo.DestinationState,
		announcementInfo.DestinationCountry,
		announcementInfo.NumberOfTravelers,
		announcementInfo.CardId)
	if err != nil {
		return err
	}
	query = `delete from announcement_language where announcement_id = $1`
	_, err = s.DB.ExecContext(ctx, query, announcementInfo.CardId)
	return err
}

func (s *PostgresRepository) DeleteAnnouncement(announcement_id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := `delete from announcement_offer where announcement_id = $1`
	_, err := s.DB.ExecContext(ctx, query, announcement_id)
	if err != nil {
		return err
	}
	query = `delete from chatlist where announcement_id = $1`
	_, err = s.DB.ExecContext(ctx, query, announcement_id)
	if err != nil {
		return err
	}
	query = `delete from announcement where id = $1`
	_, err = s.DB.ExecContext(ctx, query, announcement_id)
	return err
}

func (s *PostgresRepository) UpdatePostInformation(postInfo *pb.EditPostRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := `update post set  
		title = COALESCE(NULLIF($1, ''), title),
		rating = COALESCE(NULLIF($2, 0), rating),
		body = COALESCE(NULLIF($3, ''), body)
		where id = $4`
	_, err := s.DB.ExecContext(ctx, query,
		postInfo.PostTitle,
		postInfo.HostRating,
		postInfo.PostBody,
		postInfo.PostId)
	return err
}

func (s *PostgresRepository) GetHostId(announcement_id int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	var host_id int
	query := `select main_host from announcement where id = $1`
	err := s.DB.QueryRowContext(ctx, query, announcement_id).Scan(&host_id)
	if err != nil {
		return -1, err
	}
	return host_id, err
}

func (s *PostgresRepository) UpdateHostImagesCount(user_id, imageCount int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := `update members set hostImageCount=$1 where id=$2`
	_, err := s.DB.ExecContext(ctx, query, imageCount, user_id)
	return err
}

func (s *PostgresRepository) DeleteUserChatList(announcement_id, host_id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := `delete from chatlist where announcement_id = $1 and host_id = $2`
	_, err := s.DB.ExecContext(ctx, query, announcement_id, host_id)
	if err != nil {
		return err
	}
	return nil
}
func (s *PostgresRepository) UpdateChatListStatus(announcement_id, host_id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := `update chatlist set chat_status = $1 where announcement_id = $2 and host_id != $3`
	_, err := s.DB.ExecContext(ctx, query, 2, announcement_id, host_id)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresRepository) UpdateAnnouncementStatus(announcement_id, host_id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := `update announcement set announcement_status=$1 , main_host = $2 where id = $3`
	_, err := s.DB.ExecContext(ctx, query, 2, host_id, announcement_id)
	if err != nil {
		return err
	}
	return nil
}
