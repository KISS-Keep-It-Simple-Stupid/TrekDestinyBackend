package queue

import (
	"encoding/json"
	"log"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/notification/dbrepo"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/notification/models"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Queue struct {
	ConnString string
	QueueName  string
	channel    *amqp.Channel
	connection *amqp.Connection
	Hub        *models.Hub
	Repo       dbrepo.Repository
}

func (q *Queue) New() error {
	conn, err := amqp.Dial(q.ConnString)
	if err != nil {
		return err
	}
	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return err
	}
	q.channel = channel
	q.connection = conn
	return nil
}
func (q *Queue) Up() {
	defer q.connection.Close()
	defer q.channel.Close()
	msgs, err := q.channel.Consume(
		q.QueueName, // queue
		"",          // consumer
		true,        // auto ack
		false,       // exclusive
		false,       // no local
		false,       // no wait
		nil,         //args
	)
	if err != nil {
		log.Fatal(err)
	}

	for d := range msgs {
		temp := models.NotifMessage{}
		err := json.Unmarshal(d.Body, &temp)
		if err != nil {
			log.Println(err)
		}
		go q.Repo.InsertNotification(temp)
		q.Hub.Send <- temp.UserID
	}
}
