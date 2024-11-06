package config

import (
	"github.com/spf13/viper"
	"log"
)

// Config represents the application configuration.
type Config struct {
	Server ServerConfig `yaml:"server"`
	Redis  RedisConfig  `yaml:"redis"`
}

// ServerConfig represents the server configuration.
type ServerConfig struct {
	Name   string `yaml:"name"`
	Port   string `yaml:"port"`
	Locale string `yaml:"locale"`
}

// RedisConfig represents the Redis configuration.
type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Addr     string `yaml:"addr"`
	DB       int    `yaml:"db"`
	Password string `yaml:"password"`
}

func InitConfig() (*Config, error) {
	// 设置配置文件名和路径
	viper.SetConfigName("config")   // 配置文件名（不带扩展名）
	viper.SetConfigType("yaml")     // 如果配置文件名中没有扩展名，则需要设置类型
	viper.AddConfigPath("./config") // 查找配置文件的路径

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// 获取并打印配置文件路径
	configFilePath := viper.ConfigFileUsed()
	log.Printf("Using config file: %s", configFilePath)

	// 解析配置文件到结构体
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return &config, nil
}
