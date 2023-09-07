package utils

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/yazilimcigenclik/dream-ai-backend/models"
)

func SendDreamToQueue(dream *models.Dream) {
	var conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		fmt.Println("error connecting to rabbitmq", err)
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		fmt.Println("error opening channel", err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"new-dream",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println("error declaring queue", err)
	}

	body := fmt.Sprintf(`{"id": %d}`, dream.ID)

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

	if err != nil {
		fmt.Println("error publishing message", err)
	}

	fmt.Println("message published successfully")
}
