package apis

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type TimeResponse struct {
	Result    string `json:"result"`
	ServiceId string `json:"serviceId"`
}

func QueryTimeByStyle(style, addr string, port int) string {
	url := fmt.Sprintf("http://%s:%d/api/getDateTime", addr, port)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("http.NewRequest failed, err:", err)
		return ""
	}
	req.Header.Set("style", style)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client.Do failed, err:", err)
		return ""
	}
	rspBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body: ", err)
		return ""
	}
	log.Println("Response status code:", resp.Status)
	log.Println("Response body:", string(rspBody))
	var timeRsp TimeResponse
	err = json.Unmarshal(rspBody, &timeRsp)
	if err != nil {
		log.Println("Error unmarshalling response body: ", err)
		return ""
	}
	return timeRsp.Result
}
