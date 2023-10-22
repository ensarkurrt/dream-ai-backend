package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yazilimcigenclik/dream-ai-backend/app/constants"
	"github.com/yazilimcigenclik/dream-ai-backend/app/domain/dto"
	"github.com/yazilimcigenclik/dream-ai-backend/app/pkg"
	"github.com/yazilimcigenclik/dream-ai-backend/app/services"
	"net/http"
	"strconv"
)

type DreamController interface {
	GetAllDreams(c *gin.Context)
	CreateDream(c *gin.Context)
	GetDreamById(c *gin.Context)
}

type DreamControllerImpl struct {
	svc services.DreamService
}

func (d *DreamControllerImpl) GetAllDreams(c *gin.Context) {
	defer pkg.PanicHandler(c)

	userId, exists := c.Get("user_id")

	if exists != true {
		pkg.PanicException(constants.InvalidRequest)
	}
	dreams := d.svc.GetAllDream(userId.(uint))

	c.JSON(http.StatusOK, pkg.BuildResponse[[]dto.DreamDTO](constants.Success, dreams))
}

func (d *DreamControllerImpl) CreateDream(c *gin.Context) {
	defer pkg.PanicHandler(c)
	var request dto.CreateDreamRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		pkg.PanicException(constants.InvalidRequest)
	}

	userId, exists := c.Get("user_id")

	if exists != true {
		pkg.PanicException(constants.InvalidRequest)
	}
	request.UserID = userId.(uint)
	dream := d.svc.CreateDream(request)

	c.JSON(http.StatusOK, pkg.BuildResponse[dto.DreamDTO](constants.Success, dream))
}

func (d *DreamControllerImpl) GetDreamById(c *gin.Context) {
	defer pkg.PanicHandler(c)
	dreamID, err := strconv.Atoi(c.Param("id"))

	if dreamID == 0 || err != nil {
		pkg.PanicException(constants.InvalidRequest)
	}

	dream := d.svc.GetDreamById(dreamID)

	c.JSON(http.StatusOK, pkg.BuildResponse[dto.DreamDTO](constants.Success, dream))
}

func NewDreamController(dreamService services.DreamService) *DreamControllerImpl {
	return &DreamControllerImpl{
		svc: dreamService,
	}
}
