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
	AuthCtrl  controllers.AuthController
}

func NewInitialization(
	db *gorm.DB,
) *Initialization {
	log.Println("Initialization started")

	dreamQueueRepository := repository.NewDreamQueueRepository(db)
	userRepository := repository.NewUserRepository(db)
	dreamRepository := repository.NewDreamRepository(db)

	dreamService := services.NewDreamService(dreamRepository, dreamQueueRepository, userRepository)
	dreamController := controllers.NewDreamController(dreamService)

	userService := services.NewUserService(userRepository)
	authController := controllers.NewAuthController(userService)

	return &Initialization{
		DreamCtrl: dreamController,
		AuthCtrl:  authController,
	}

}
