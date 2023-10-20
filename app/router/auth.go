package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yazilimcigenclik/dream-ai-backend/config"
)

func NewAuthRoutes(init *config.Initialization, group *gin.RouterGroup) {
	authGroup := group.Group("/auth")

	authGroup.POST("/login", init.AuthCtrl.Login)
	authGroup.POST("/register", init.AuthCtrl.Register)
}
