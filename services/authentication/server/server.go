package server

import (
	"context"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/authentication/db"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/authentication/email"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/authentication/helper"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/authentication/models"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/authentication/pb"
	"github.com/spf13/viper"
)

type Repository struct {
	pb.UnimplementedAuthServer
	DB db.Repository
}

func New(db db.Repository) *Repository {
	return &Repository{
		DB: db,
	}
}

func (s *Repository) Signup(ctx context.Context, r *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	check, err := s.DB.CheckUserExistance(r.Email, r.UserName)
	if err != nil {
		respErr := errors.New("internal server error while checking user existance - authentication service")
		log.Println(err)
		return nil, respErr
	}
	if check {
		resp := pb.SignUpResponse{
			Message: "user already exists",
		}
		return &resp, nil
	}
	err = s.DB.InsertUser(r)
	if err != nil {
		respErr := errors.New("internal server error while adding new user - authentication service")
		log.Println(err)
		return nil, respErr
	}

	// sending verification email
	var (
		from             = viper.Get("EMAILHOST").(string)
		password         = viper.Get("EMAILPASSWORD").(string)
		frontend_address = viper.Get("FRONTEND_ADDRESS").(string)
	)
	access_exp_time, _ := strconv.Atoi(viper.Get("ACCESS_EXP_TIME").(string))
	// create jwt token
	jwtClaim := models.JwtClaims{
		UserName: r.UserName,
		ExpDate:  time.Now().Add(time.Duration(time.Duration(access_exp_time) * time.Minute)),
	}
	token, err := helper.NewToken(&jwtClaim)
	if err != nil {
		respErr := errors.New("internal server error while generating jwt token for verification - authentication service")
		log.Println(err)
		return nil, respErr
	}

	// creating email
	verificationEmail := email.Email{
		From:     from,
		Password: password,
		To:       []string{r.Email},
		Text:     "Hello " + r.FirstName + "\ncheck link below to verify your email\n" + frontend_address + "/verify-email?token=" + token,
	}

	go verificationEmail.Send()

	resp := pb.SignUpResponse{
		Message: "success",
	}
	return &resp, nil
}

func (s *Repository) Login(ctx context.Context, r *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, nil
}
