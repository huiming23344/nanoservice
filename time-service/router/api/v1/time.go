package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/huiming23344/nanoservice/time-service/server"
	"time"
)

type TimeResponse struct {
	Result    string `json:"result"`
	ServiceId string `json:"serviceId"`
}

func QueryTime(c *gin.Context) {
	style := c.GetHeader("style")
	var timeStr string
	switch style {
	case "full":
		timeStr = time.Now().UTC().Format("2006-01-02 15:04:05")
	case "date":
		timeStr = time.Now().UTC().Format("2006-01-02")
	case "time":
		timeStr = time.Now().UTC().Format("15:04:05")
	case "unix":
		timeStr = fmt.Sprint(time.Now().UTC().UnixMilli())
	}
	rsp := TimeResponse{
		Result:    timeStr,
		ServiceId: server.TimeServer.ServiceId,
	}
	c.JSON(200, rsp)
}
