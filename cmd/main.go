package main

import (
	"apibaffle/config"
	"apibaffle/config/redis"
	"apibaffle/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// 加载配置
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	// 初始化 Redis 客户端
	rdb, err := redis.NewRedisClient(&conf.Redis)
	// 设置 Gin 模式
	gin.SetMode(gin.ReleaseMode)
	// 创建 Gin 路由实例
	r := gin.Default()
	// 设置路由
	router.InitRouter(r, rdb)
	// 启动服务器
	serverPort := ":" + conf.Server.Port
	log.Printf("Starting server on port %s", serverPort)
	err1 := r.Run(serverPort)
	if err1 != nil {
		log.Fatalf("Error starting server: %v", err1)
	}
}
