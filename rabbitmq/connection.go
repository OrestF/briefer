package rabbitmq

import (
	"github.com/gobuffalo/envy"
	"github.com/streadway/amqp"
)

var CONN *amqp.Connection

func Connect() (conn *amqp.Connection) {
	url := envy.Get("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/")
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")

	CONN = conn
	return conn
}

func CloseConnection() {
	CONN.Close()
}
