package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/yazilimcigenclik/dream-ai-backend/models"
	"github.com/yazilimcigenclik/dream-ai-backend/utils"
	"net/http"
)

func GetAllDreams(c *gin.Context) {
	var dreams []models.Dream
	models.DB.Find(&dreams)

	utils.RespondWithJSON(c, http.StatusOK, dreams)
}

func GetDream(c *gin.Context) {

	id := c.Param("id")
	var dream models.Dream

	if err := models.DB.Where("id = ?", id).First(&dream).Error; err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Dream not found!")
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, dream)

}

var validate *validator.Validate

func CreateDream(c *gin.Context) {
	var input models.DreamCreateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Validation Error. Please check your inputs")
		return
	}

	dream := &models.Dream{
		Content: input.Content,
	}

	result := models.DB.Create(dream)

	if result.Error != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "An error occurred while creating dream")
		return
	}

	if result.RowsAffected == 0 {
		utils.RespondWithError(c, http.StatusInternalServerError, "An error occurred while creating dream")
		return
	}

	utils.RespondWithJSON(c, http.StatusCreated, dream)

}
