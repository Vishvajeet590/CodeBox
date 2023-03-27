package rabbitMQ

import (
	"github.com/rabbitmq/amqp091-go"
	"os"
)

func getExchangeName() string {
	return os.Getenv("RMQ_EXCHANGE_NAME")
}
func getQueueName() string {
	return os.Getenv("RMQ_QUEUE_NAME")
}

func declareRandomQueue(ch *amqp091.Channel) (amqp091.Queue, error) {
	q, err := ch.QueueInspect(getQueueName())
	if err != nil {
		// Queue not exist
		return ch.QueueDeclare(
			getQueueName(), // name
			false,          // durable
			false,          // delete when unused
			true,           // exclusive
			false,          // no-wait
			nil,            // arguments
		)
	}
	return q, nil
}

func declareExchange(ch *amqp091.Channel) error {
	return ch.ExchangeDeclare(
		getExchangeName(), // name
		"topic",           // type
		true,              // durable
		false,             // auto-deleted
		false,             // internal
		false,             // no-wait
		nil,               // arguments
	)
}
