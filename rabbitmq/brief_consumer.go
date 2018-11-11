package rabbitmq

import (
	"log"
)

func StartProjectsBriefConsumer() {
	consumer := InitConsumer()
	queueName := "projects.brief"
	consumer.defineQueue(queueName)

	msgs, err := consumer.channel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	failOnError(err, "Failed to register a consumer")

	briefForever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("AMQP Received a message from %s: %s", queueName, d.Body)
			client := InitClient()
			client.Publish("callbacks.projects.brief", string(d.Body))
		}
	}()

	log.Printf(" [*] Waiting for messages from %s. To exit press CTRL+C", queueName)
	<-briefForever
}
