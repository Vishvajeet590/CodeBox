package rabbitMQ

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"time"
)

// Emitter for publishing AMQP events
type Emitter struct {
	connection *amqp091.Connection
}

func (e *Emitter) setup() error {
	channel, err := e.connection.Channel()
	if err != nil {
		panic(err)
	}

	_, err = channel.QueueDeclare(
		os.Getenv("RMQ_QUEUE_NAME"), // name
		true,                        // durable
		false,                       // delete when unused
		false,                       // exclusive
		false,                       // no-wait
		nil,                         // arguments
	)

	defer channel.Close()
	return declareExchange(channel)
}

// Push (Publish) a specified message to the AMQP exchange
func (e *Emitter) Push(event string, severity string) error {
	channel, err := e.connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = channel.PublishWithContext(
		ctx,
		getExchangeName(),
		severity,
		false,
		false,
		amqp091.Publishing{
			DeliveryMode: amqp091.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(event),
		},
	)
	log.Printf("Sending message: %s -> %s", event, getExchangeName())
	return nil
}

func NewEventEmitter(conn *amqp091.Connection) (Emitter, error) {
	emitter := Emitter{
		connection: conn,
	}

	err := emitter.setup()
	if err != nil {
		return Emitter{}, err
	}

	return emitter, nil
}
