# api-baffle
## 项目概述
实际项目中存在上游接口时间段关闭、接口收费、无足够测试数据案例的情况

## 目的
用于模拟构建虚假接口（类似mock，但实际项目情况复杂）
因为之前在公司开发的项目挡板属于公司财产（java版），因此用Go语言和相关技术栈构建简单复现大致框架

## 需要的软件和工具

在开始之前，请确保你已经安装了以下软件和工具：

- [Go](https://golang.org/)（版本 1.18 或更高）
- [Docker](https://www.docker.com/)
- [Redis](https://redis.io/)
- [Gin](https://gin-gonic.com/)
- [Viper](https://github.com/spf13/viper)
- [Redis Client for Go](https://github.com/go-redis/redis/v9)

## 安装

1. **克隆仓库：**

   ```bash
   git clone https://github.com/yourusername/data-shield-system.git
   cd data-shield-system
2. **安装依赖：**
    ```bash
    go mod download
3. **配置 Redis：**
确保 Redis 在你的本地机器或远程服务器上运行。你可以使用 Docker 启动 Redis：
   ```bash
    docker run -d --name redis -p 6379:6379 redis
4. **设置配置文件：**
复制 config.yaml.example 文件为 config.yaml 并根据你的环境进行修改：
    ```bash
    server:
    port: 8080

    redis:
    addr: "localhost:6379"
    password: ""
    db: 0
4. **运行应用程序：**
    ```bash
    go run main.go
**访问 API**： API 将在 http://localhost:8080 上可用。
测试 API
你可以使用 curl 或 Postman 等工具来测试 API 端点。

如果有任何问题或需要进一步的帮助，请联系项目维护者。