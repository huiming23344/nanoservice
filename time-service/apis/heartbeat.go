package apis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/huiming23344/nanoservice/registry"
	"io"
	"log"
	"net/http"
)

type HBRequest struct {
	ServiceId string `json:"serviceId"`
	IpAddress string `json:"ipAddress"`
	Port      string `json:"port"`
}

func Heartbeat() {
	body := HBRequest{
		ServiceId: registry.TimeServer.ServiceId,
		IpAddress: registry.TimeServer.IpAddress,
		Port:      registry.TimeServer.Port,
	}
	jsonData, _ := json.Marshal(body)
	reqBody := bytes.NewBuffer(jsonData)
	url := fmt.Sprintf("http://&s:%s/api/heartbeat", registry.TimeServer.Registry.Address, registry.TimeServer.Registry.Port)
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		fmt.Println("http.NewRequest failed, err:", err)
		return
	}
	client := &http.Client{}
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
