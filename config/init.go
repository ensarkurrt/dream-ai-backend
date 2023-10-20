package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/yazilimcigenclik/dream-ai-backend/app/controllers"
	"github.com/yazilimcigenclik/dream-ai-backend/app/repository"
	"github.com/yazilimcigenclik/dream-ai-backend/app/services"
	"gorm.io/gorm"
)

type Initialization struct {
	DreamCtrl controllers.DreamController
}

func NewInitialization(
	db *gorm.DB,
) *Initialization {
	log.Println("Initialization started")
	dreamRepository := repository.DreamRepositoryInit(db)
	dreamService := services.DreamServiceInit(dreamRepository)
	dreamController := controllers.DreamControllerInit(dreamService)

	return &Initialization{
		DreamCtrl: dreamController,
	}

}
