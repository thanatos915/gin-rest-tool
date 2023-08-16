package cmd

import (
	"gin-rest-tool/app/router"
	"gin-rest-tool/global"
	"log"
	"net/http"
)

// 运行服务
func RunServer() {
	// 运行服务
	s := &http.Server{
		Addr:           ":" + global.C.Server.HttpPort,
		Handler:        router.InitRouter(),
		ReadTimeout:    global.C.Server.ReadTimeout,
		WriteTimeout:   global.C.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("服务启动失败: %v", err)
	}
}
