package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yazilimcigenclik/dream-ai-backend/config"
)

func Init(init *config.Initialization) *gin.Engine {
	router := gin.New()
	log.Println("router init")
	router.Use(gin.Logger())
	log.Println("router init 2")
	router.Use(gin.Recovery())

	log.Println("router init 3")
	// Dream Controller Endpoints
	/*router.GET("/dreams", GetAllDreams)
	router.GET("/dream/:id", GetDream)
	router.POST("/dream", CreateDream)*/

	api := router.Group("/api")
	{
		log.Println("router group")
		dream := api.Group("/dream")
		log.Println("router group 2")
		dream.GET("", init.DreamCtrl.GetAllDreams)
		dream.POST("", init.DreamCtrl.CreateDream)
		dream.GET("/:id", init.DreamCtrl.GetDreamById)
	}

	return router
}
