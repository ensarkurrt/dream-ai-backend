package services

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/yazilimcigenclik/dream-ai-backend/app/constants"
	"github.com/yazilimcigenclik/dream-ai-backend/app/domain/dao"
	"github.com/yazilimcigenclik/dream-ai-backend/app/pkg"
	"github.com/yazilimcigenclik/dream-ai-backend/app/repository"
)

type DreamService interface {
	GetAllDream(c *gin.Context)
	GetDreamById(c *gin.Context)
	CreateDream(c *gin.Context)
}

type DreamServiceImpl struct {
	dreamRepository repository.DreamRepository
}

func (u DreamServiceImpl) GetDreamById(c *gin.Context) {
	defer pkg.PanicHandler(c)

	dreamID, _ := strconv.Atoi(c.Param("id"))

	data, err := u.dreamRepository.FindDreamById(dreamID)

	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constants.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constants.Success, data))
}

func (u DreamServiceImpl) GetAllDream(c *gin.Context) {
	defer pkg.PanicHandler(c)

	dreams, err := u.dreamRepository.FindAllDream()

	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constants.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constants.Success, dreams))
}

func (u DreamServiceImpl) CreateDream(c *gin.Context) {
	defer pkg.PanicHandler(c)

	var dream dao.Dream

	if err := c.ShouldBindJSON(&dream); err != nil {
		fmt.Println("Error occurred while binding dream on model", err)
		pkg.PanicException(constants.InvalidRequest)
	}

	data, err := u.dreamRepository.CreateDream(&dream)

	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constants.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constants.Success, data))
}

func DreamServiceInit(dreamRepository repository.DreamRepository) *DreamServiceImpl {
	return &DreamServiceImpl{
		dreamRepository: dreamRepository,
	}
}
