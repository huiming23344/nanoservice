package registry

import (
	"fmt"
	"github.com/huiming23344/nanoservice/registry/config"
	routers "github.com/huiming23344/nanoservice/registry/router"
	"github.com/huiming23344/nanoservice/registry/server"
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

	router := routers.InitRouter()

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: router,
	}

	server.InitRegistryServer()
	if err := s.ListenAndServe(); err != nil {
		log.Printf("Listen: %s\n", err)
	}
}
