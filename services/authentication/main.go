package main

import (
	"fmt"
	"log"
	"net"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/authentication/db"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/authentication/driver"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/authentication/pb"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/authentication/server"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
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
	s := grpc.NewServer()
	pb.RegisterAuthServer(s, server)
	fmt.Printf("authentication grpc server is listening on port %s\n", port)
	s.Serve(lis)
}
