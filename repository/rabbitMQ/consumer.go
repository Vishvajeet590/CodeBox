package rabbitMQ

import (
	"CodeBox/judge/runner"
	"CodeBox/models"
	"context"
	"encoding/json"
	"github.com/rabbitmq/amqp091-go"
	"golang.org/x/sync/errgroup"
	"log"
)

type Consumer struct {
	conn      *amqp091.Connection
	queueName string
}

func (consumer *Consumer) setup() error {
	channel, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	return declareExchange(channel)
}

// NewConsumer returns a new Consumer
func NewConsumer(conn *amqp091.Connection) (Consumer, error) {
	consumer := Consumer{
		conn: conn,
	}
	err := consumer.setup()
	if err != nil {
		return Consumer{}, err
	}

	return consumer, nil
}

// Listen will listen for all new Queue publications
// and print them to the console.
func (consumer *Consumer) Listen(topics []string, runner *runner.Runner) error {
	ch, err := consumer.conn.Channel()
	ch.Qos(3, 0, false)

	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := declareRandomQueue(ch)
	if err != nil {
		return err
	}

	for _, s := range topics {
		err = ch.QueueBind(
			q.Name,
			s,
			getExchangeName(),
			false,
			nil,
		)

		if err != nil {
			return err
		}
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	errs, _ := errgroup.WithContext(context.Background())

	for d := range msgs {

		d := d
		errs.Go(
			func() error {
				var config models.TaskPackage
				err = json.Unmarshal(d.Body, &config)
				if err != nil {
					return err
				}

				err = runner.Run(config)
				//	time.Sleep(time.Duration(rand.IntnRange(100, 1000)) * time.Second)
				if err != nil {
					return err
				}

				return nil
			})

	}
	return errs.Wait()
	/*go func() {

	}()*/

	log.Printf("[*] Waiting for message [Exchange, Queue][%s, %s]. To exit press CTRL+C", getExchangeName(), q.Name)
	<-forever
	return nil
}
