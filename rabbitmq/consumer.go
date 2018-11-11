package rabbitmq

import (
	"log"

	"github.com/gobuffalo/envy"
	"github.com/streadway/amqp"
)

func InitConsumer() (consumer Consumer) {
	consumer = Consumer{connection: *CONN}
	consumer.DefineChannel()

	return consumer
}

type Consumer struct {
	connection amqp.Connection
	channel    amqp.Channel
}

func (c *Consumer) Connect() {
	url := envy.Get("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/")
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")

	c.connection = *conn
}

func (c *Consumer) DefineChannel() {
	ch, err := c.connection.Channel()
	failOnError(err, "Failed to open a channel")

	c.channel = *ch
}

func (c *Consumer) defineQueue(queueName string) (q amqp.Queue) {
	q, err := c.channel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	return q
}

func (c *Consumer) Listen(queueName string) {
	c.defineQueue(queueName)

	msgs, err := c.channel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("AMQP Received a message from %s: %s", queueName, d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages from %s. To exit press CTRL+C", queueName)
	<-forever
}
