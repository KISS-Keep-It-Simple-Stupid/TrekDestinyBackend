package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/userprofile/helper"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/userprofile/pb"
	"github.com/aws/aws-sdk-go/service/s3"
	"golang.org/x/crypto/bcrypt"
)

type PostgresRepository struct {
	DB *sql.DB
}

func NewPostgresRepository(db *sql.DB) Repository {
	return &PostgresRepository{
		DB: db,
	}
}

func (s *PostgresRepository) GetUserDetails(username string) (*pb.ProfileDetailsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	user := pb.ProfileDetailsResponse{}
	var birth_date, joiningdate time.Time
	query := `select email, username, firstname, lastname, birthdate, state, country, gender, joiningdate, ishost, 
			COALESCE(NULLIF(bio, NULL), '') as bio,
			COALESCE(NULLIF(city, NULL), '') as city,
			COALESCE(NULLIF(address, NULL), '') as address,
			COALESCE(NULLIF(phonenumber, NULL), '') as phonenumber,
			COALESCE(NULLIF(ispetfirendly::text, ''), '') as ispetfirendly,
			COALESCE(NULLIF(iskidfiendly::text, ''), '') as iskidfiendly,
			COALESCE(NULLIF(issmokingallowed::text, ''), '') as issmokingallowed,
			COALESCE(NULLIF(roomnumber, NULL), 0) as roomnumber
			from members where username = $1`
	err := s.DB.QueryRowContext(ctx, query, username).Scan(
		&user.Email,
		&user.UserName,
		&user.FirstName,
		&user.LastName,
		&birth_date,
		&user.State,
		&user.Country,
		&user.Gender,
		&joiningdate,
		&user.IsHost,
		&user.Bio,
		&user.City,
		&user.Address,
		&user.PhoneNumber,
		&user.IsPetFriendly,
		&user.IsKidFriendly,
		&user.IsSmokingAllowed,
		&user.RoomNumber)
	if err != nil {
		return nil, err
	}
	user.BirthDate = birth_date.Format("2006-01-02")
	user.JoiningDate = joiningdate.Format("2006-01-02")
	return &user, nil
}

func (s *PostgresRepository) UpdateUserInformation(username string, userInfo *pb.EditProfileRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	password := ""
	if userInfo.CurrentPassword != "" {
		hashed_password, err := bcrypt.GenerateFromPassword([]byte(userInfo.NewPassword), bcrypt.MinCost)
		if err != nil {
			return err
		}
		password = string(hashed_password)
	}
	query := `update members set  
		password = COALESCE(NULLIF($1, ''), password),
		firstname = COALESCE(NULLIF($2, ''), firstname),
		lastname = COALESCE(NULLIF($3, ''), lastname),
		city = COALESCE(NULLIF($4, ''), city),
		state = COALESCE(NULLIF($5, ''), state),
		country = COALESCE(NULLIF($6, ''), country),
		bio = COALESCE(NULLIF($7, ''), bio),
		ishost = COALESCE(NULLIF($8, '')::boolean, ishost),
		address = COALESCE(NULLIF($9, ''), address),
		ispetfirendly = COALESCE(NULLIF($10, '')::boolean, ispetfirendly),
		iskidfiendly = COALESCE(NULLIF($11, '')::boolean, iskidfiendly),
		issmokingallowed = COALESCE(NULLIF($12, '')::boolean, issmokingallowed),
		phonenumber = COALESCE(NULLIF($13, ''), phonenumber),
		roomnumber = COALESCE(NULLIF($14, 0), roomnumber)
		where username = $15`
	_, err := s.DB.ExecContext(ctx, query,
		password,
		userInfo.FirstName,
		userInfo.LastName,
		userInfo.City,
		userInfo.State,
		userInfo.Country,
		userInfo.Bio,
		userInfo.IsHost,
		userInfo.Address,
		userInfo.IsPetFriendly,
		userInfo.IsKidFriendly,
		userInfo.IsSmokingAllowed,
		userInfo.PhoneNumber,
		userInfo.RoomNumber,
		username)
	return err
}

func (s *PostgresRepository) CheckUserExistance(userUserName string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	temp := 0
	query := `select id from members where username = $1`
	err := s.DB.QueryRowContext(ctx, query, userUserName).Scan(&temp)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}

		return false, nil
	}
	return true, nil
}

func (s *PostgresRepository) GetUserPassword(username string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	current_password := ""
	query := `select password from members where username = $1`
	err := s.DB.QueryRowContext(ctx, query, username).Scan(&current_password)
	if err != nil {
		return "", err
	}
	return current_password, nil
}

func (s *PostgresRepository) GetPublicProfile(username string) (*pb.PublicProfileResponse, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	user := pb.PublicProfileResponse{}
	var birth_date, joiningdate time.Time
	var id int
	query := `select id ,email, username, firstname, lastname, birthdate, state, country, gender, joiningdate, 
			COALESCE(NULLIF(bio, NULL), '') as bio,
			COALESCE(NULLIF(city, NULL), '') as city
			from members where username = $1`
	err := s.DB.QueryRowContext(ctx, query, username).Scan(
		&id,
		&user.Email,
		&user.UserName,
		&user.FirstName,
		&user.LastName,
		&birth_date,
		&user.State,
		&user.Country,
		&user.Gender,
		&joiningdate,
		&user.Bio,
		&user.City)
	if err != nil {
		return nil, 0, err
	}
	user.BirthDate = birth_date.Format("2006-01-02")
	user.JoiningDate = joiningdate.Format("2006-01-02")
	return &user, id, nil
}

func (s *PostgresRepository) GetPublicProfileHost(guest_id int, host_username string) (*pb.PublicProfileHostResponse, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	user := pb.PublicProfileHostResponse{}
	host_id, err := s.GetIdFromUsername(host_username)
	if err != nil {
		return nil, "", err
	}
	check, err := s.CheckIfUserCanViewProfile(guest_id, host_id)
	if err != nil {
		return nil, "", err
	}
	if !check {
		message := "you can't view this profile because the user hasn't offered to any of your announcements"
		return &user, message, nil
	}
	var birth_date, joiningdate time.Time
	query := `select email, username, firstname, lastname, birthdate, state, country, gender, joiningdate, ishost, 
			COALESCE(NULLIF(bio, NULL), '') as bio,
			COALESCE(NULLIF(city, NULL), '') as city,
			COALESCE(NULLIF(address, NULL), '') as address,
			COALESCE(NULLIF(phonenumber, NULL), '') as phonenumber,
			COALESCE(NULLIF(ispetfirendly::text, ''), '') as ispetfirendly,
			COALESCE(NULLIF(iskidfiendly::text, ''), '') as iskidfiendly,
			COALESCE(NULLIF(issmokingallowed::text, ''), '') as issmokingallowed,
			COALESCE(NULLIF(roomnumber, NULL), 0) as roomnumber
			from members where username = $1`
	err = s.DB.QueryRowContext(ctx, query, host_username).Scan(
		&user.Email,
		&user.UserName,
		&user.FirstName,
		&user.LastName,
		&birth_date,
		&user.State,
		&user.Country,
		&user.Gender,
		&joiningdate,
		&user.IsHost,
		&user.Bio,
		&user.City,
		&user.Address,
		&user.PhoneNumber,
		&user.IsPetFriendly,
		&user.IsKidFriendly,
		&user.IsSmokingAllowed,
		&user.RoomNumber)
	if err != nil {
		return nil, "", err
	}
	user.BirthDate = birth_date.Format("2006-01-02")
	user.JoiningDate = joiningdate.Format("2006-01-02")
	return &user, "", nil
}

func (s *PostgresRepository) CheckIfUserCanViewProfile(guest_id, host_id int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := `select announcement_id from announcement_offer where host_id = $1`
	rows, err := s.DB.QueryContext(ctx, query, host_id)
	if err != nil {
		return false, err
	}
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return false, err
		}
		var user_id int
		query := `select user_id from announcement where id = $1 and user_id = $2`
		err = s.DB.QueryRowContext(ctx, query, id, guest_id).Scan(&user_id)
		if err != nil {
			if err != sql.ErrNoRows {
				return false, err
			}
		}
		if user_id == guest_id {
			return true, nil
		}
	}
	return false, nil
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

func (s *PostgresRepository) InsertChatList(host_id, guest_id, announcement_id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := `insert into chatlist (host_id , guest_id , announcement_id )values ($1 , $2 , $3 )`
	_, err := s.DB.ExecContext(ctx, query, host_id, guest_id, announcement_id)
	return err
}

func (s *PostgresRepository) GetChatList(guest_id int, obj *s3.S3) (*pb.ChatListResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := `select id ,host_id , guest_id, announcement_id , chat_status from chatlist where guest_id = $1 or host_id = $2`
	rows, err := s.DB.QueryContext(ctx, query, guest_id, guest_id)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
		return &pb.ChatListResponse{
			Users: make([]*pb.ChatList, 0),
		}, nil
	}

	users := make([]*pb.ChatList, 0)
	for rows.Next() {
		temp_res := &pb.ChatList{}
		temp1 := 0
		temp2 := 0
		err := rows.Scan(&temp_res.ID, &temp1, &temp2, &temp_res.AnnoucementId, &temp_res.Status)
		if err != nil {
			return nil, err
		}
		if temp1 == guest_id {
			temp_res.IsHost = "yes"
			temp_res.HostID = int32(temp2)
			temp_res.Username, err = s.GetUserNameByID(temp2)
			if err != nil {
				return nil, err
			}
		} else {
			temp_res.IsHost = "no"
			temp_res.HostID = int32(temp1)
			temp_res.Username, err = s.GetUserNameByID(temp1)
			if err != nil {
				return nil, err
			}
		}
		temp_res.Image, err = helper.GetImageURL(obj, fmt.Sprintf("user-%d", temp_res.HostID))
		if err != nil {
			return nil, err
		}
		users = append(users, temp_res)
	}

	resp := &pb.ChatListResponse{
		Users: users,
	}

	return resp, nil
}

func (s *PostgresRepository) GetUserNameByID(user_id int) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := `select username from members where id = $1`
	username := ""
	err := s.DB.QueryRowContext(ctx, query, user_id).Scan(&username)
	return username, err
}

func (s *PostgresRepository) GetHouseImagesCount(user_id int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := `select hostImageCount from members where id = $1`
	count := 0
	err := s.DB.QueryRowContext(ctx, query, user_id).Scan(&count)
	return count, err
}

func (s *PostgresRepository) UpdateOfferStatus(announcement_id, host_id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := `update announcement_offer set offer_status=$1 where announcement_id = $2 and host_id = $3`
	_, err := s.DB.ExecContext(ctx, query, 2, announcement_id, host_id)
	return err
}

func (s *PostgresRepository) InsertUserLanguage(user_id int, lang string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := `insert into member_language (user_id, language) values ($1, $2)`
	_, err := s.DB.ExecContext(ctx, query, user_id, lang)
	return err
}

func (s *PostgresRepository) InsertUserInterest(user_id int, interest string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := `insert into member_interest (user_id, interest) values ($1, $2)`
	_, err := s.DB.ExecContext(ctx, query, user_id, interest)
	return err
}