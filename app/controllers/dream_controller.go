package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yazilimcigenclik/dream-ai-backend/app/services"
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
	d.svc.GetAllDream(c)
}

func (d *DreamControllerImpl) CreateDream(c *gin.Context) {
	d.svc.CreateDream(c)
}

func (d *DreamControllerImpl) GetDreamById(c *gin.Context) {
	d.svc.GetDreamById(c)
}

func DreamControllerInit(dreamService services.DreamService) *DreamControllerImpl {
	return &DreamControllerImpl{
		svc: dreamService,
	}
}
