package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RespondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"error": message})
}

func RespondWithJSON(c *gin.Context, code int, message string, payload interface{}) {

	var responseMap gin.H

	if payload == nil {
		responseMap = gin.H{
			"message": message,
		}
	} else {
		responseMap = gin.H{
			"message": message,
			"data":    payload,
		}
	}

	c.JSON(code, responseMap)
}

func formatJSON(data []byte) string {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", " ")

	if err != nil {
		fmt.Println(err)
	}

	d := out.Bytes()
	return string(d)
}
