package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Publish(exchange string, routingKey string, msg string) error {
	_, err := connectRabbit()
	if err != nil {
		return err
	}

	// err = channel.PublishWithContext(context.Background(), exchange, routingKey, false, false, amqp.Publishing{
	// 	ContentType: "text/plain",
	// 	Body:        []byte(msg),
	// })

	// if err != nil {
	// 	return err
	// }

	log.Fatalf("nao deu error")
	return nil
}

func Sub() {

}

func connectRabbit() (*amqp.Channel, error) {
	con, err := amqp.Dial("amqp://admin:admin123@localhost:5672/")
	if err != nil {
		return nil, err
	}
	channel, err := con.Channel()
	if err != nil {
		return nil, err
	}

	return channel, nil
}
