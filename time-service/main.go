package main

import (
	"fmt"
	"github.com/huiming23344/nanoservice/time-service/apis"
	"github.com/huiming23344/nanoservice/time-service/config"
	routers "github.com/huiming23344/nanoservice/time-service/router"
	"github.com/huiming23344/nanoservice/time-service/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Load config failed: %v", err)
	}
	config.SetGlobalConfig(cfg)

	// 创建一个信号通道
	sigChan := make(chan os.Signal, 1)

	// 注册信号通道，以接收 os.Interrupt 和 syscall.SIGTERM
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// 启动一个 goroutine 来监听信号
	go func() {
		<-sigChan // 等待接收信号
		fmt.Println("Received an interrupt, cleaning up...")

		// 这里放置清理逻辑
		// 例如：关闭文件描述符、断开网络连接等

		fmt.Println("Cleanup finished, exiting...")
		os.Exit(0) // 退出程序
	}()

	router := routers.InitRouter()

	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Server.Port),
		Handler: router,
	}

	server.InitTimeServer()
	go apis.Heartbeat()
	if err := s.ListenAndServe(); err != nil {
		log.Printf("Listen: %s\n", err)
	}
}
