package server

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/huiming23344/nanoservice/log-service/config"
	"net"
)

type Registry struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

type logServer struct {
	ServiceName string
	ServiceId   string
	IpAddress   string
	Port        int
	Registry    Registry
	LogList     []*Log
}

type Log struct {
	ServiceName string `json:"serviceName"`
	ServiceId   string `json:"serviceId"`
	Datetime    string `json:"datetime"`
	Level       string `json:"level"`
	Message     string `json:"message"`
}

var LogServer logServer

func InitTimeServer() {
	cfg := config.GlobalConfig()
	serviceId := uuid.New().String()
	addrs, err := getHostIPAddresses()
	if err != nil {
		fmt.Printf("get host ip failed: %v\n", err)
		return
	}
	fmt.Printf("serviceId: %s\n", serviceId)
	LogServer = logServer{
		ServiceName: "log-service",
		ServiceId:   serviceId,
		IpAddress:   addrs[len(addrs)-1],
		Port:        cfg.Server.Port,
		Registry: Registry{
			Address: cfg.Registry.Address,
			Port:    cfg.Registry.Port,
		},
	}
}

func SetLog(log *Log) {
	LogServer.LogList = append(LogServer.LogList, log)
}

func getHostIPAddresses() ([]string, error) {
	var addresses []string

	// 获取所有网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	// 遍历网络接口
	for _, i := range interfaces {
		// 获取接口的地址列表
		addrs, err := i.Addrs()
		if err != nil {
			return nil, err
		}

		// 遍历地址列表
		for _, addr := range addrs {
			// 检查是否为IPv4地址
			ip := addr.(*net.IPNet)
			if ip.IP.To4() != nil {
				// 添加IPv4地址到结果列表
				addresses = append(addresses, ip.IP.String())
			}
		}
	}

	return addresses, nil
}
