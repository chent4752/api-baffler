package router

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

// DataRequest 结构体用于解析请求体中的数据
type DataRequest struct {
	ID   string                 `json:"id"`
	Data map[string]interface{} `json:"data"`
}

func putDataHandler(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req DataRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 创建带有超时的上下文
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// 将 Data 转换为 JSON 字符串
		dataJSON, err := json.Marshal(req.Data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data"})
			return
		}

		if err := rdb.Set(ctx, req.ID, dataJSON, 0).Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Data created or updated"})
	}
}

func deleteDataHandler(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		// 验证输入的合法性
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		// 使用请求上下文
		ctx := c.Request.Context()
		if err := rdb.Del(ctx, id).Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Data deleted"})
	}
}

func getDataHandler(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		// 创建带有超时的上下文
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		data, err := rdb.Get(ctx, id).Result()
		if err == redis.Nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 将 JSON 字符串转换回 map
		var dataMap map[string]interface{}
		if err := json.Unmarshal([]byte(data), &dataMap); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal data"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": dataMap})
	}
}
