package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/authentication/pb"
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

func (s *PostgresRepository) InsertUser(user *pb.SignUpRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := `insert into member (email , username , password , firstname , lastname , birthdate , city , country)
			  values ($1,$2,$3,$4,$5,$6,$7,$8)`
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	birthdate, err := time.Parse("2006-01-02", user.BirthDate)
	if err != nil {
		return err
	}
	_, err = s.DB.ExecContext(ctx, query, user.Email, user.UserName, string(hashedPassword),
		user.FirstName, user.LastName, birthdate, user.City, user.Country)
	return err
}

func (s *PostgresRepository) CheckUserExistance(userEmail, userUserName string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	temp := 0
	query := `select id from member where email = $1 or username = $2`
	err := s.DB.QueryRowContext(ctx, query, userEmail, userUserName).Scan(&temp)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}

		return false, nil
	}
	return true, nil
}
