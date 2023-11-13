package main

import (
	"fmt"
	"log"
	"net"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/userprofile/db"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/userprofile/driver"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/userprofile/pb"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/userprofile/server"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	// setup config variables
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	var (
		port       = viper.Get("SERVERPORT").(string)
		region     = viper.Get("OBJECT_STORAGE_REGION").(string)
		access_key = viper.Get("OBJECT_STORAGE_ACCESS_KEY").(string)
		secret_key = viper.Get("OBJECT_STORAGE_SECRET_KEY").(string)
		endpoint   = viper.Get("OBJECT_STORAGE_ENDPOINT").(string)
	)

	// setup database
	dbConn, err := driver.NewDBConnection()
	if err != nil {
		log.Fatalln(err.Error())
	}
	// setup object storage
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String(region),
		Endpoint: aws.String(endpoint),
		Credentials: credentials.NewStaticCredentials(
			access_key,
			secret_key,
			"",
		),
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	// Create an S3 client
	svc := s3.New(sess)

	repository := db.NewPostgresRepository(dbConn)
	server := server.New(repository, svc)
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// setting up grpc server
	s := grpc.NewServer()
	pb.RegisterUserProfileServer(s, server)
	fmt.Printf("userprofile grpc server is listening on port %s\n", port)
	s.Serve(lis)
}
