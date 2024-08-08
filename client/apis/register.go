package apis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/huiming23344/nanoservice/client/server"
	"io"
	"log"
	"net/http"
)

type Service struct {
	ServiceName string `json:"serviceName"`
	ServiceId   string `json:"serviceId"`
	IpAddress   string `json:"ipAddress"`
	Port        int    `json:"port"`
}

func Register() {
	body := Service{
		ServiceName: server.ClientServer.ServiceName,
		ServiceId:   server.ClientServer.ServiceId,
		IpAddress:   server.ClientServer.IpAddress,
		Port:        server.ClientServer.Port,
	}
	jsonData, _ := json.Marshal(body)
	reqBody := bytes.NewBuffer(jsonData)
	url := fmt.Sprintf("http://%s:%d/api/register", server.ClientServer.Registry.Address, server.ClientServer.Registry.Port)
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		fmt.Println("http.NewRequest failed, err:", err)
		return
	}
	client := &http.Client{}
	log.Println("Heartbeat to registry...")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client.Do failed, err:", err)
		return
	}
	repBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body: ", err)
		return
	}
	log.Println("Response status code:", resp.Status)
	log.Println("Response body:", string(repBody))
}

func Unregister() {
	body := Service{
		ServiceName: server.ClientServer.ServiceName,
		ServiceId:   server.ClientServer.ServiceId,
		IpAddress:   server.ClientServer.IpAddress,
		Port:        server.ClientServer.Port,
	}
	jsonData, _ := json.Marshal(body)
	reqBody := bytes.NewBuffer(jsonData)
	url := fmt.Sprintf("http://%s:d/api/unregister", server.ClientServer.Registry.Address, server.ClientServer.Registry.Port)
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		fmt.Println("http.NewRequest failed, err:", err)
		return
	}
	client := &http.Client{}
	log.Println("Heartbeat to registry...")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client.Do failed, err:", err)
		return
	}
	repBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body: ", err)
		return
	}
	log.Println("Response status code:", resp.Status)
	log.Println("Response body:", string(repBody))
}
