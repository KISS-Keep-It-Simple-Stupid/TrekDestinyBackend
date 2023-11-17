package main

import (
	"fmt"
	"log"
	"net"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/announcement/db"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/announcement/driver"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/announcement/pb"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/announcement/queue"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/announcement/server"
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
		port            = viper.Get("SERVERPORT").(string)
		rabbit_user     = viper.Get("RABBIT_USER").(string)
		rabbit_password = viper.Get("RABBIT_PASSWORD").(string)
		rabbit_port     = viper.Get("RABBIT_PORT").(string)
		rabbit_host     = viper.Get("RABBIT_HOST").(string)
		queue_name      = viper.Get("QUEUENAME").(string)
		region          = viper.Get("OBJECT_STORAGE_REGION").(string)
		access_key      = viper.Get("OBJECT_STORAGE_ACCESS_KEY").(string)
		secret_key      = viper.Get("OBJECT_STORAGE_SECRET_KEY").(string)
		endpoint        = viper.Get("OBJECT_STORAGE_ENDPOINT").(string)
	)

	// setup database
	dbConn, err := driver.NewDBConnection()
	if err != nil {
		log.Fatalln(err.Error())
	}

	// initialize rabbitMQ
	rabbit_conn := fmt.Sprintf("amqp://%s:%s@%s:%s", rabbit_password, rabbit_user, rabbit_host, rabbit_port)
	rabbitQueue := &queue.Queue{
		ConnString: rabbit_conn,
		QueueName:  queue_name,
	}
	err = rabbitQueue.New()
	if err != nil {
		log.Fatal(err)
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
	server := server.New(repository, rabbitQueue, svc)
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// setting up grpc server
	s := grpc.NewServer()
	pb.RegisterAnnouncementServer(s, server)
	fmt.Printf("announcement grpc server is listening on port %s\n", port)
	s.Serve(lis)
}
