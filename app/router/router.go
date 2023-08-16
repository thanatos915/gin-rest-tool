package router

import (
	v1 "gin-rest-tool/app/controller/v1"
	"gin-rest-tool/app/middleware"
	"gin-rest-tool/docs"
	"gin-rest-tool/global"
	"gin-rest-tool/pkg/limiter"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title API接口文档
// @version 1.0.o
// @description 这是一个API接口文档
// @BasePath /api/v1
func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.ContextTimeout(global.C.Server.ContextTimeout))
	if global.C.Server.RunMode != "debug" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		// 设置swagger
		docs.SwaggerInfo.BasePath = "/api/v1"
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	// 健康检查
	r.GET("/api/v1/ping", v1.Ping)

	// 速率限制
	var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
		limiter.LimiterBucketRule{
			Key:          "/debug/vars",
			FillInterval: time.Second,
			Capacity:     20000,
			Quantum:      20000,
		},
	)
	r.Use(middleware.RateLimiter(methodLimiters))

	return r
}
