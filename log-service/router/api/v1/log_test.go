package v1

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
)

func TestQueryTime(t *testing.T) {
	queryTimeByStyle("full")
	queryTimeByStyle("date")
	queryTimeByStyle("time")
	queryTimeByStyle("unix")
}

func queryTimeByStyle(style string) {
	url := "http://127.0.0.1:8280/api/getDateTime"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("http.NewRequest failed, err:", err)
		return
	}
	req.Header.Set("style", style)
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
