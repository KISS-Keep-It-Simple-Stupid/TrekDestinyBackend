package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/userprofile/pb"
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
	query := `select email,username,firstname ,lastname , birthdate , city , country , gender , bio , state , joiningdate from members where username = $1`
	err := s.DB.QueryRowContext(ctx, query, username).Scan(
		&user.Email,
		&user.UserName,
		&user.FirstName,
		&user.LastName,
		&birth_date,
		&user.City,
		&user.Country,
		&user.Gender,
		&user.Bio,
		&user.State,
		&joiningdate)
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
	password = COALESCE(NULLIF($1, ''), password) ,
	firstname = COALESCE(NULLIF($2, ''), firstname),
	lastname = COALESCE(NULLIF($3, ''), lastname),
	city = COALESCE(NULLIF($4, ''), city),
	country = COALESCE(NULLIF($5, ''), country),
	bio = COALESCE(NULLIF($6, ''), bio)
	where username=$7`
	_, err := s.DB.ExecContext(ctx, query,
		password,
		userInfo.FirstName,
		userInfo.LastName,
		userInfo.City,
		userInfo.Country,
		userInfo.Bio,
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
