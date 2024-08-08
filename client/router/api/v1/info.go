package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/huiming23344/nanoservice/client/apis"
	"github.com/huiming23344/nanoservice/client/server"
	"time"
)

type TimeResp struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}

func QueryInfo(c *gin.Context) {
	services := apis.Discovery("time-service")
	if len(services) == 0 {
		c.JSON(200, TimeResp{
			Error:  "No available time-service",
			Result: "",
		})
		return
	}
	service := services[0]
	timeRsp := apis.QueryTimeByStyle("full", service.IpAddress, service.Port)
	timeGet, err := time.Parse(timeRsp, "2006-01-02 15:04:05")
	if err != nil {
		c.JSON(200, TimeResp{
			Error:  "Parse time failed",
			Result: "",
		})
		return
	}
	timeLocal := timeGet.Local().Format("2006-01-02 15:04:05")
	result := fmt.Sprintf("Hello Kingsoft Cloud Star Camp - %s - %s", server.ClientServer.ServiceId, timeLocal)
	c.JSON(200, TimeResp{
		Error:  "",
		Result: result,
	})
}
