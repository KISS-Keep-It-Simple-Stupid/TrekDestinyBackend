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
	repository := db.NewPostgresRepository(dbConn)
	server := server.New(repository, rabbitQueue)
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
