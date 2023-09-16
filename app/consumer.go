package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/yazilimcigenclik/dream-ai-backend/models"
	"github.com/yazilimcigenclik/dream-ai-backend/utils"
	"log"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	models.ConnectDatabase()

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()

	_, err = ch.QueueDeclare(
		"new-dream",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println("error declaring queue", err)
		return
	}

	msgs, err := ch.Consume(
		"new-dream",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var data dreamQueue
			err := json.Unmarshal([]byte(d.Body), &data)
			if err != nil {
				fmt.Println("error decoding response body", err)
				return
			}

			fmt.Println(data.DreamId)

			go createDream(&data.DreamId)
		}
	}()

	println("Consumer started")

	<-forever
}

func createDream(dreamId *uint) {

	var dream models.Dream

	var dreamQueue models.DreamImageQueue

	fmt.Println("Dream ID: ", *dreamId)

	if err := models.DB.Where("id = ?", *dreamId).First(&dream).Error; err != nil {
		fmt.Println("Dream not found!")
		return
	}

	explanationChan := make(chan string)
	titleChan := make(chan string)
	imageStatusChan := make(chan string)

	go func() {
		_exp, err := utils.GenerateExplanation(dream.Content)

		if err != nil {
			fmt.Println("An error occurred while responding to your request")
			explanationChan <- ""
			return
		}

		fmt.Println("Explanation: ", *_exp)

		explanationChan <- *_exp
	}()

	go func() {
		_title, err := utils.GenerateTitle(dream.Content)
		if err != nil {
			fmt.Println("An error occurred while responding to your request")
			titleChan <- ""
			return
		}

		fmt.Println("Title: ", *_title)
		titleChan <- *_title
	}()

	explanation := <-explanationChan
	title := <-titleChan

	models.DB.Model(&dream).Updates(models.Dream{
		Explanation: explanation,
		Title:       title,
	})

	if err := models.DB.Where("dream_id = ?", dreamId).First(&dreamQueue).Error; err == nil && (dreamQueue.Status != "succeeded" && dreamQueue.Status != "failed") {
		fmt.Println("Image generation is already in progress")
		go getImageStatus(dreamQueue, imageStatusChan)

		for {
			select {
			case <-imageStatusChan:
				break
			}
		}
	}

}

func getImageStatus(dreamQueue models.DreamImageQueue, imageStatusChan chan string) {

	queue, err := utils.UpdateStatusFromAPI(dreamQueue)
	if err != nil {
		fmt.Println("An error occurred while responding to your request for image")
		return
	}

	fmt.Println("Image generation started")

	if queue.Status == "succeeded" || queue.Status == "failed" {
		fmt.Println("Image generation completed")
		imageStatusChan <- queue.Status
	}

}

type dreamQueue struct {
	DreamId uint `json:"id"`
}