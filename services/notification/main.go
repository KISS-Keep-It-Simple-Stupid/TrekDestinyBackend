package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/notification/dbrepo"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/notification/driver"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/notification/handlers"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/notification/hub"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/notification/queue"
	"github.com/spf13/viper"
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
	myhub := hub.New()
	go hub.Up(myhub)
	db, err := driver.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	dbRepo := dbrepo.New(db)

	// initilize rabbitMQ
	rabbit_conn := fmt.Sprintf("amqp://%s:%s@%s:%s", rabbit_password, rabbit_user, rabbit_host, rabbit_port)
	rabbitQueue := &queue.Queue{
		ConnString: rabbit_conn,
		QueueName:  queue_name,
		Hub:        myhub,
		Repo:       dbRepo,
	}
	err = rabbitQueue.New()
	if err != nil {
		log.Fatal(err)
	}
	handler := handlers.New(myhub, dbRepo)
	go rabbitQueue.Up()
	server := http.Server{
		Addr:    ":" + port,
		Handler: getRoutes(handler),
	}
	fmt.Println("notification service listens on port", port)
	server.ListenAndServe()
}
