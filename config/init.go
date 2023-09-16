package config

import (
	"log"

	"github.com/yazilimcigenclik/dream-ai-backend/app/controllers"
	"github.com/yazilimcigenclik/dream-ai-backend/app/repository"
	"github.com/yazilimcigenclik/dream-ai-backend/app/services"
)

type Initialization struct {
	DreamRepo repository.DreamRepository
	DreamSvc  services.DreamService
	DreamCtrl controllers.DreamController
}

func NewInitialization(
	dreamRepo repository.DreamRepository,
	dreamService services.DreamService,
	dreamCtrl controllers.DreamController,
) *Initialization {
	log.Print("Initializing...")
	return &Initialization{
		DreamRepo: dreamRepo,
		DreamSvc:  dreamService,
		DreamCtrl: dreamCtrl,
	}
}
