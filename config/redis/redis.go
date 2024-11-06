package redis

import (
	"apibaffle/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

// NewRedisClient 根据提供的配置创建并初始化一个新的 Redis 客户端。
// 参数:
//
//	cfg: 指向 config.RedisConfig 的指针，包含连接到 Redis 所需的配置信息。
//
// 返回值:
//
//	*redis.Client: 初始化后的 Redis 客户端指针。
//	error: 如果连接失败则返回错误。
func NewRedisClient(cfg *config.RedisConfig) (*redis.Client, error) {
	// 使用给定的配置初始化一个新的 Redis 客户端。
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	// 测试连接
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		// 如果连接测试失败，返回错误。
		return nil, fmt.Errorf("连接到 Redis 失败: %w", err)
	}

	// 记录成功初始化的日志信息。
	log.Println("Redis 客户端初始化成功")
	return client, nil
}
