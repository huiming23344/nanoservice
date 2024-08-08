package server

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/huiming23344/nanoservice/time-service/config"
)

type Registry struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}

type timeServer struct {
	ServiceName string
	ServiceId   string
	IpAddress   string
	Port        string
	Registry    Registry
}

var TimeServer timeServer

func InitTimeServer() {
	cfg := config.GlobalConfig()
	serviceId := uuid.New().String()
	fmt.Printf("serviceId: %s\n", serviceId)
	TimeServer = timeServer{
		ServiceName: "time-service",
		ServiceId:   serviceId,
		IpAddress:   "",
		Port:        cfg.Server.Port,
		Registry: Registry{
			Address: cfg.Registry.Address,
			Port:    cfg.Registry.Port,
		},
	}
}
