package rabbitmq

import (
	"log"

	"github.com/gobuffalo/envy"
	"github.com/streadway/amqp"
)

func InitClient() (client Client) {
	client = Client{connection: *CONN}
	client.DefineChannel()

	return client
}

type Client struct {
	connection amqp.Connection
	channel    amqp.Channel
}

func (c *Client) Connect() {
	url := envy.Get("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/")
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")

	c.connection = *conn
}

func (c *Client) CloseConnection() {
	c.connection.Close()
}

func (c *Client) DefineChannel() {
	ch, err := c.connection.Channel()
	failOnError(err, "Failed to open a channel")

	c.channel = *ch
}

func (c *Client) Publish(queueName string, message string) {
	queue := c.defineQueue(queueName)

	err := c.channel.Publish(
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	failOnError(err, "Failed to publish a message")
}

func (c *Client) PublishAsync(queueName string, message string) {
	go func(client *Client, queueName string, message string) {
		client.Publish(queueName, message)
	}(c, queueName, message)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (c *Client) defineQueue(queueName string) (q amqp.Queue) {
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
