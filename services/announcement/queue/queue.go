package queue

import (
	"context"
	"encoding/json"
	"time"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/announcement/models"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Queue struct {
	ConnString string
	QueueName  string
	channel    *amqp.Channel
	connection *amqp.Connection
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
	_, err = channel.QueueDeclare(
		q.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

func (q *Queue) Send(message *models.NotificationMessage) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Millisecond)
	defer cancel()
	notfiBody, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return q.channel.PublishWithContext(ctx,
		"",
		q.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "appliction/json",
			Body:        notfiBody,
		})
}
