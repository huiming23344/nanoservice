package v1

import (
	"github.com/gin-gonic/gin"
)

type LogReq struct {
	ServiceName string `json:"serviceName"`
	ServiceId   string `json:"serviceId"`
	Datetime    string `json:"datetime"`
	Level       string `json:"level"`
	Message     string `json:"message"`
}

func SetLog(c *gin.Context) {
	var logReq LogReq
	if err := c.ShouldBindJSON(&logReq); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}
