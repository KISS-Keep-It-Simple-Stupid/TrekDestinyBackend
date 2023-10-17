package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/gateway/handlers"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	var (
		auth_service_address = viper.Get("AUTH_SERVICE_ADDRESS").(string)
		port                 = viper.Get("SERVERPORT").(string)
	)
	auth_conn, err := grpc.Dial(auth_service_address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err.Error())
	}
	handler := handlers.New(auth_conn)
	routes := getRoutes(handler)
	server := http.Server{
		Addr:    ":" + port,
		Handler: routes,
	}

	fmt.Println("gateway is up on port :", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
