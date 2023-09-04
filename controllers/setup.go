package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
)

func New() {
	router := gin.Default()

	// Dream Controller Endpoints
	router.GET("/dreams", GetAllDreams)
	router.GET("/dream/:id", GetDream)
	router.POST("/dream", CreateDream)

	log.Fatal(router.Run(":8008"))
}
