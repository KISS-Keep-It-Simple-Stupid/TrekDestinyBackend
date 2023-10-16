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
	"golang.org/x/crypto/bcrypt"
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
	user, check, err := s.DB.GetLoginCridentials(r.Email)

	// check user existance
	if err != nil {
		respErr := errors.New("internal server error while checking user existance - authentication service")
		log.Println(err)
		return nil, respErr
	}
	if !check {
		resp := pb.LoginResponse{
			Message: "There is no registered user corresponding to the given email",
		}
		return &resp, nil
	}

	// checking verification
	if !user.IsVerified {
		// sending verification email
		var (
			from             = viper.Get("EMAILHOST").(string)
			password         = viper.Get("EMAILPASSWORD").(string)
			frontend_address = viper.Get("FRONTEND_ADDRESS").(string)
		)
		access_exp_time, _ := strconv.Atoi(viper.Get("ACCESS_EXP_TIME").(string))
		// create jwt token
		jwtClaim := models.JwtClaims{
			UserName: user.UserName,
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
			To:       []string{user.Email},
			Text:     "Hello " + user.FirstName + "\ncheck link below to verify your email\n" + frontend_address + "/verify-email?token=" + token,
		}

		go verificationEmail.Send()
		resp := pb.LoginResponse{
			Message: "User is not verified",
		}
		return &resp, nil
	}

	// checking cridentials
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password))
	if err != nil {
		resp := &pb.LoginResponse{
			Message: "invalid password",
		}
		return resp, nil
	}

	// valid user - now can login

	access_exp_time, _ := strconv.Atoi(viper.Get("ACCESS_EXP_TIME").(string))
	refresh_exp_time, _ := strconv.Atoi(viper.Get("REFRESH_EXP_TIME").(string))
	access_claims := &models.JwtClaims{
		UserName: user.UserName,
		ExpDate:  time.Now().Add(time.Duration(time.Duration(access_exp_time) * time.Minute)),
	}
	refresh_claims := &models.JwtClaims{
		UserName: user.UserName,
		ExpDate:  time.Now().Add(time.Duration(time.Duration(refresh_exp_time) * time.Hour * 24)),
	}

	refresh, err := helper.NewToken(refresh_claims)
	if err != nil {
		respErr := errors.New("internal server error while generating jwt refresh token for login - authentication service")
		log.Println(err)
		return nil, respErr
	}
	access, err := helper.NewToken(access_claims)
	if err != nil {
		respErr := errors.New("internal server error while generating jwt access token for login - authentication service")
		log.Println(err)
		return nil, respErr
	}

	loginResp := &pb.LoginResponse{
		Message:      "success",
		AccessToken:  access,
		RefreshToken: refresh,
	}

	return loginResp, nil
}
