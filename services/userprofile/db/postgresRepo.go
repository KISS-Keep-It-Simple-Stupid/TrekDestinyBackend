package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/userprofile/pb"
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
	var birth_date time.Time
	query := `select email,username,firstname ,lastname , birthdate , city , country , gender , bio from members where username = $1`
	err := s.DB.QueryRowContext(ctx, query, username).Scan(
		&user.Email,
		&user.UserName,
		&user.FirstName,
		&user.LastName,
		&birth_date,
		&user.City,
		&user.Country,
		&user.Gender,
		&user.Bio)
	if err != nil {
		return nil, err
	}
	user.BirthDate = birth_date.Format("2006-01-02")
	return &user, nil
}
