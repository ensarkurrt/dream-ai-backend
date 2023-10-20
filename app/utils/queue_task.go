package utils

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/yazilimcigenclik/dream-ai-backend/app/domain/dao"
	"os"
)

func NewDreamQueueTask(dreamQueue dao.DreamQueue) error {

	mq_url := os.Getenv("RABBITMQ_CONNECTION")
	log.Println("RabbitMQ bağlantısı kuruluyor. Env: ", mq_url)
	// RabbitMQ bağlantısı kur
	conn, err := amqp.Dial(mq_url)
	if err != nil {
		log.Fatalf("RabbitMQ bağlantısı kurulamadı: %s", err)
		return err
	}
	defer conn.Close()

	// Kanal oluştur
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Kanal oluşturulamadı: %s", err)
		return err
	}
	defer ch.Close()

	// İş sırasını tanımla
	queueName := "dream_queue"

	// DreamQueue struct'ını JSON formatına dönüştür
	dreamJSON, err := json.Marshal(dreamQueue)
	if err != nil {
		log.Fatalf("DreamQueue struct'ı JSON'a dönüştürülemedi: %s", err)
		return err
	}

	// Mesajı gönder
	err = ch.Publish("", queueName, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        dreamJSON,
	})

	if err != nil {
		log.Fatalf("Mesaj gönderilemedi: %s", err)
		return err
	}

	log.Println("Mesaj başarıyla gönderildi")
	return nil
}
