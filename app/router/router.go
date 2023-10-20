package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yazilimcigenclik/dream-ai-backend/app/router/middlewares"
	"github.com/yazilimcigenclik/dream-ai-backend/config"
)

func Init(init *config.Initialization) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		NewAuthRoutes(init, api)

		v1 := api.Group("/v1")
		{
			v1.Use(middlewares.AuthMiddleware())
			NewDreamRouters(init, v1)
		}
	}

	return router
}
