package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/chat/db"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/chat/driver"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/chat/handlers"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/chat/socket/hub"
	"github.com/spf13/viper"
)

func main() {

	// setup config variables
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	port := viper.Get("SERVERPORT").(string)
	// setup database
	dbConn, err := driver.NewDBConnection()
	if err != nil {
		log.Fatalln(err.Error())
	}
	DbRepo := db.NewPostgresRepository(dbConn)
	myhub := hub.NewHub()
	go hub.StartHub(myhub)
	handler := handlers.NewHandler(myhub, DbRepo)
	server := http.Server{
		Addr:    ":" + port,
		Handler: GetRoutes(handler),
	}
	fmt.Println("Chat service is running on port:", port)
	server.ListenAndServe()
}
