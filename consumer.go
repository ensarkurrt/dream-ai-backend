package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/yazilimcigenclik/dream-ai-backend/app/domain/dao"
	"github.com/yazilimcigenclik/dream-ai-backend/app/repository"
	"github.com/yazilimcigenclik/dream-ai-backend/app/utils"
	"github.com/yazilimcigenclik/dream-ai-backend/config"
	"gorm.io/gorm"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.InitLog()

	db := config.ConnectToDB()

	dreamRepository := repository.NewDreamRepository(db)
	dreamQueueRepository := repository.NewDreamQueueRepository(db)

	// RabbitMQ bağlantısı kur
	log.Println("RabbitMQ bağlantısı kuruluyor. Env: ", os.Getenv("RABBITMQ_CONNECTION"))
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_CONNECTION"))
	if err != nil {
		log.Fatalf("RabbitMQ bağlantısı kurulamadı: %s", err)
	}

	consumer := NewConsumer(
		db,
		dreamRepository,
		dreamQueueRepository,
		conn)

	consumer.Consume()
}

type Consumer interface {
	ProcessDreamQueue(dreamQueue *dao.DreamQueue) error
	Consume()
}

type ConsumerImpl struct {
	db                   *gorm.DB
	dreamRepository      repository.DreamRepository
	dreamQueueRepository repository.DreamQueueRepository
	connection           *amqp.Connection
}

func (consumer *ConsumerImpl) processDreamQueue(dreamQueue *dao.DreamQueue) error {
	gpt := utils.NewGPT()
	dreamQueue.Status = dao.Processing
	dreamQueue.Dream.Status = dao.Processing

	_, err := consumer.dreamQueueRepository.Update(*dreamQueue)
	_, err = consumer.dreamRepository.UpdateDream(dreamQueue.Dream)

	explanation, err := gpt.GenerateExplanation(dreamQueue.Dream.Content)
	if err != nil {
		log.Error("Happened error when generating Explanation: Error: ", err)
		dreamQueue.Status = dao.Failed
		dreamQueue.Dream.Status = dao.Failed
		_, err = consumer.dreamQueueRepository.Update(*dreamQueue)
		_, err = consumer.dreamRepository.UpdateDream(dreamQueue.Dream)
		return err
	}

	title, err := gpt.GenerateTitle(dreamQueue.Dream.Content)
	if err != nil {
		log.Error("Happened error when generating Title: Error: ", err)
		dreamQueue.Status = dao.Failed
		dreamQueue.Dream.Status = dao.Failed
		_, err = consumer.dreamQueueRepository.Update(*dreamQueue)
		_, err = consumer.dreamRepository.UpdateDream(dreamQueue.Dream)
		return err
	}

	dreamQueue.Dream.Explanation = explanation
	dreamQueue.Dream.Title = title
	dreamQueue.Dream.Status = dao.Completed
	dreamQueue.Status = dao.Completed

	_, err = consumer.dreamQueueRepository.Update(*dreamQueue)
	if err != nil {
		log.Error("Happened error when updating dream queue: Error: ", err)
		dreamQueue.Status = dao.Failed
		dreamQueue.Dream.Status = dao.Failed
		_, err = consumer.dreamQueueRepository.Update(*dreamQueue)
		_, err = consumer.dreamRepository.UpdateDream(dreamQueue.Dream)
		return err
	}

	_, err = consumer.dreamRepository.UpdateDream(dreamQueue.Dream)
	if err != nil {
		log.Error("Happened error when updating dream: Error: ", err)
		dreamQueue.Status = dao.Failed
		dreamQueue.Dream.Status = dao.Failed
		_, err = consumer.dreamQueueRepository.Update(*dreamQueue)
		_, err = consumer.dreamRepository.UpdateDream(dreamQueue.Dream)
		return err
	}

	return nil
}

func (consumer *ConsumerImpl) Consume() {

	queueName := "dream_queue"
	defer consumer.connection.Close()

	// Kanal oluştur
	ch, err := consumer.connection.Channel()
	if err != nil {
		log.Fatalf("Kanal oluşturulamadı: %s", err)
	}
	defer ch.Close()

	// İş sırasını tanımla
	msgs, err := ch.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("İş sırası dinlenemedi: %s", err)
	}
	forever := make(chan bool)
	// Mesajları işle
	for msg := range msgs {
		log.Printf("Mesaj alındı: %s", msg.Body)

		// Mesajı JSON'dan DreamQueue struct'ına çöz
		var dreamQueue dao.DreamQueue
		err := json.Unmarshal(msg.Body, &dreamQueue)
		if err != nil {
			log.Printf("Mesaj çözümlenemedi: %s", err)
			continue
		}

		log.Printf("Mesaj İçeriği: %+v", dreamQueue)

		err = consumer.processDreamQueue(&dreamQueue)
		if err != nil {
			log.Printf("Mesaj işlenemedi: %s", err)
			continue
		}

		err = msg.Ack(false)
		if err != nil {
			log.Printf("Mesaj işlendi olarak işaretlenemedi: %s", err)
			continue
		}
	}

	<-forever
}

func NewConsumer(db *gorm.DB, dreamRepository repository.DreamRepository, dreamQueueRepository repository.DreamQueueRepository, con *amqp.Connection) *ConsumerImpl {

	return &ConsumerImpl{
		db,
		dreamRepository,
		dreamQueueRepository,
		con,
	}
}
