package rabbitmq

import (
	"encoding/json"
	"log"
)

func StartProjectsBriefConsumer() {
	consumer := InitConsumer()
	queueName := "estimator.projects.brief"
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
			pp := ParsedProject{}
			json.Unmarshal([]byte(d.Body), &pp)
			producer := ProjectProducer{Id: pp.Id, Title: pp.Title, Description: pp.Description}
			producer.Push()
		}
	}()

	log.Printf(" [*] Waiting for messages from %s. To exit press CTRL+C", queueName)
	<-briefForever
}

type ParsedProject struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
