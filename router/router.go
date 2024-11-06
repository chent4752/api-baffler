package router

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func InitRouter(r *gin.Engine, rdb *redis.Client) {

	r.PUT("/data", putDataHandler(rdb))
	r.DELETE("/data/:id", deleteDataHandler(rdb))
	r.GET("/data/:id", getDataHandler(rdb))
}
