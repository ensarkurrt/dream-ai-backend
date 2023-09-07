package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/yazilimcigenclik/dream-ai-backend/models"
	"net/http"
	"os"
)

func GenerateImageWithPrompt(dream models.Dream) (*models.DreamImageQueue, error) {

	apiUrl := "https://api.replicate.com/v1/predictions"
	apiToken := os.Getenv("REPLICATE_API_TOKEN")
	postData := []byte(`{
    "version": "ad59ca21177f9e217b9075e7300cf6e14f7e5b4505b87b9689dbd866e9768969",
    "input": {
        "prompt": "mdjrny-v4 style. ` + dream.Content + `., digital painting, concept art, smooth, sharp focus, illustration, 8k",
        "width": 512,
        "height": 512,
        "guidance_scale": 7,
        "num_inference_steps": 50,
        "num_outputs": 1,
        "prompt_strength": 0.8,
        "scheduler": "KLMS"
    }
}`)

	// create new http request
	request, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(postData))
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Token "+apiToken)

	// send the request
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err)
	}

	var imageResponse ImageResponse

	err = json.NewDecoder(response.Body).Decode(&imageResponse)

	if err != nil {
		fmt.Println("error decoding response body", err)
		return nil, err
	}

	// clean up memory after execution
	defer response.Body.Close()

	if imageResponse.Error != nil {
		return nil, fmt.Errorf(*imageResponse.Error)
	}

	dreamQueue := models.DreamImageQueue{
		DreamId: dream.ID,
		QueueId: imageResponse.Id,
		Version: imageResponse.Version,
		GetUrl:  imageResponse.Urls.Get,
		Status:  imageResponse.Status,
	}

	models.DB.Create(&dreamQueue)

	return &dreamQueue, nil
}

func UpdateStatusFromAPI(imageRequest models.DreamImageQueue) (*models.DreamImageQueue, error) {

	apiToken := os.Getenv("REPLICATE_API_TOKEN")
	// create new http request
	request, err := http.NewRequest("GET", imageRequest.GetUrl, nil)
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Token "+apiToken)

	// send the request
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err)
	}

	var imageResponse ImageResponse

	err = json.NewDecoder(response.Body).Decode(&imageResponse)

	if err != nil {
		fmt.Println("error decoding response body", err)
		return nil, err
	}

	// clean up memory after execution
	defer response.Body.Close()

	if imageResponse.Error != nil {
		return nil, fmt.Errorf(*imageResponse.Error)
	}

	if imageResponse.Status == "succeeded" {
		models.DB.Model(&imageRequest).Updates(models.DreamImageQueue{
			Output: (*imageResponse.Outputs)[0],
			Status: imageResponse.Status,
		})

		models.DB.Model(&models.Dream{}).Where("id = ?", imageRequest.DreamId).Updates(models.Dream{
			ImageUrl: &(*imageResponse.Outputs)[0],
		})

	} else {
		models.DB.Model(&imageRequest).Updates(models.DreamImageQueue{
			Status: imageResponse.Status,
		})
	}

	return &imageRequest, nil
}

type ImageResponse struct {
	Id      string    `json:"id"`
	Version string    `json:"version"`
	Outputs *[]string `json:"output"`
	Error   *string   `json:"error"`
	Status  string    `json:"status"`
	Urls    Urls      `json:"urls"`
}

type Urls struct {
	Cancel string `json:"cancel"`
	Get    string `json:"get"`
}
