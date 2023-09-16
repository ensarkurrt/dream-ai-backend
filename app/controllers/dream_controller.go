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

func (u DreamControllerImpl) GetAllDreamData(c *gin.Context) {
	u.svc.GetAllDream(c)
}

func (u DreamControllerImpl) AddDreamData(c *gin.Context) {
	u.svc.CreateDream(c)
}

func (u DreamControllerImpl) GetDreamById(c *gin.Context) {
	u.svc.GetDreamById(c)
}

func DreamControllerInit(dreamService services.DreamService) *DreamControllerImpl {
	return &DreamControllerImpl{
		svc: dreamService,
	}
}
