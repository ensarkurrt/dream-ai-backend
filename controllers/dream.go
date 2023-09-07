package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/yazilimcigenclik/dream-ai-backend/models"
	"github.com/yazilimcigenclik/dream-ai-backend/utils"
)

func GetAllDreams(c *gin.Context) {
	var dreams []models.Dream
	models.DB.Find(&dreams)

	if len(dreams) == 0 {
		utils.RespondWithError(c, http.StatusNotFound, "Dreams not found!")
		return
	}

	fmt.Println("Dreams found successfully")
	utils.RespondWithJSON(c, http.StatusOK, "Dreams found successfully", dreams)
}

func GetDream(c *gin.Context) {

	id := c.Param("id")
	var dream models.Dream

	if err := models.DB.Where("id = ?", id).First(&dream).Error; err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Dream not found!")
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, "Dream found successfully", dream)

}

var validate *validator.Validate

func CreateDream(c *gin.Context) {
	var input models.DreamCreateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Validation Error. Please check your inputs")
		return
	}

	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Validation Error. Please check your inputs")
		return
	}

	/*explanationChan := make(chan string)
	titleChan := make(chan string)

	go func() {
		_exp, err := utils.GenerateExplanation(input.Content)
		if err != nil {
			utils.RespondWithError(c, http.StatusInternalServerError, "An error occurred while responding to your request")
			explanationChan <- ""
			return
		}
		explanationChan <- *_exp
	}()

	go func() {
		_title, err := utils.GenerateTitle(input.Content)
		if err != nil {
			utils.RespondWithError(c, http.StatusInternalServerError, "An error occurred while responding to your request")
			titleChan <- ""
			return
		}
		titleChan <- *_title
	}()

	explanation := <-explanationChan
	title := <-titleChan*/

	dream := &models.Dream{
		Content: input.Content,
		/*Explanation: explanation,
		Title:       title,*/
	}

	result := models.DB.Create(dream)

	if input.GenerateImage {
		go func() {
			_, err := utils.GenerateImageWithPrompt(*dream)
			if err != nil {
				utils.RespondWithError(c, http.StatusInternalServerError, "An error occurred while responding to your request for image generation")
				return
			}
		}()
	}

	if result.Error != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "An error occurred while creating dream")
		return
	}

	if result.RowsAffected == 0 {
		utils.RespondWithError(c, http.StatusInternalServerError, "An error occurred while creating dream")
		return
	}

	utils.RespondWithJSON(c, http.StatusCreated, "Dream created successfully", nil)
}
