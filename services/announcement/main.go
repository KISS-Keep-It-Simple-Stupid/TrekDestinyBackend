package main

import (
	"fmt"
	"log"
	"net"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/announcement/db"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/announcement/driver"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/announcement/pb"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/announcement/server"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	// setup config variables
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	// setup database
	dbConn, err := driver.NewDBConnection()
	if err != nil {
		log.Fatalln(err.Error())
	}
	port := viper.Get("SERVERPORT").(string)
	repository := db.NewPostgresRepository(dbConn)
	server := server.New(repository)
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