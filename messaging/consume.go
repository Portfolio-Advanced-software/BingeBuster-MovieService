package messaging

import (
	"fmt"
	"log"

	env "github.com/Portfolio-Advanced-software/BingeBuster-MovieService/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ConsumeMessage(queue string) {

	// Retrieve values from environment variables
	rabbitmqUser := env.GoDotEnvVariable("RABBITMQ_USER")
	rabbitmqPwd := env.GoDotEnvVariable("RABBITMQ_PWD")

	// Construct the RabbitMQ URL
	rabbitmqURL := fmt.Sprintf("amqps://%s:%s@rattlesnake.rmq.cloudamqp.com/%s", rabbitmqUser, rabbitmqPwd, rabbitmqUser)

	conn, err := amqp.Dial(rabbitmqURL)
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
