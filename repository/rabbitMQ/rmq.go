package rabbitMQ

import (
	"CodeBox/judge/runner"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"os"
)

var RmqEmitter Emitter

func NewRmq() error {
	conStr := os.Getenv("RMQ_URL")
	conn, err := amqp091.Dial(conStr)
	if err != nil {
		return err
	}

	RmqEmitter, err = NewEventEmitter(conn)
	if err != nil {
		return err
	}
	return nil

}

func NewListenRmq() error {
	queues := []string{os.Getenv("RMQ_QUEUE_NAME")}
	connection, err := amqp091.Dial(os.Getenv("RMQ_URL"))
	if err != nil {
		return err
	}
	fmt.Printf("Ready-\n")
	defer func(connection *amqp091.Connection) {
		err := connection.Close()
		if err != nil {
			fmt.Printf("%v\n", err.Error())
		}
	}(connection)

	consumer, err := NewConsumer(connection)
	r := runner.NewRunner()
	if err != nil {
		return err
	}
	err = consumer.Listen(queues, r)
	if err != nil {
		return err
	}
	return nil
}
