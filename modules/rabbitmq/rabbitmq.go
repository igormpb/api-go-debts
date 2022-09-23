package rabbitmq

import (
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Publish(exchange string, routingKey string, msg []byte) error {
	channel, err := connectRabbit()
	if err != nil {
		return err
	}
	defer channel.Close()

	err = channel.Publish(exchange,
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})

	if err != nil {
		return err
	}

	return nil
}

func Sub() error {
	channel, err := connectRabbit()
	if err != nil {
		return err
	}
	defer channel.Close()

	msg, err := channel.Consume("email", "", false, false, false, false, nil)

	forver := make(chan bool)

	go func() {
		for d := range msg {
			log.Printf("message : %s", d.Body)
		}
	}()

	<-forver
	return nil
}

func connectRabbit() (*amqp.Channel, error) {
	uri := fmt.Sprintf("amqp://%s:%s@%s:5672", os.Getenv("RABBITMQ_DEFAULT_USER"), os.Getenv("RABBITMQ_DEFAULT_PASS"), "rabbitmq")

	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return channel, nil
}
