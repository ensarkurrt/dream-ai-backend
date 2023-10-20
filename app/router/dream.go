package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yazilimcigenclik/dream-ai-backend/config"
)

func NewDreamRouters(init *config.Initialization, group *gin.RouterGroup) {
	dreams := group.Group("/dreams")
	{
		dreams.GET("", init.DreamCtrl.GetAllDreams)
		dreams.POST("", init.DreamCtrl.CreateDream)
		dreams.GET("/:id", init.DreamCtrl.GetDreamById)
	}
}
