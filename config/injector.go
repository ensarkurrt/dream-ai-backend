package config

import (
	"github.com/google/wire"
	"github.com/yazilimcigenclik/dream-ai-backend/app/controllers"
	"github.com/yazilimcigenclik/dream-ai-backend/app/repository"
	"github.com/yazilimcigenclik/dream-ai-backend/app/services"
)

var db = wire.NewSet(ConnectToDB)

var dreamServiceSet = wire.NewSet(
	services.DreamServiceInit,
	wire.Bind(new(services.DreamService), new(*services.DreamServiceImpl)),
)

var dreamRepoSet = wire.NewSet(
	repository.DreamRepositoryInit,
	wire.Bind(new(repository.DreamRepository), new(*repository.DreamRepositoryImpl)),
)

var dreamCtrlSet = wire.NewSet(
	controllers.DreamControllerInit,
	wire.Bind(new(controllers.DreamController), new(*controllers.DreamControllerImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, db, dreamCtrlSet, dreamServiceSet, dreamRepoSet)
	return nil
}
