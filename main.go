package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建 Gin 引擎
	r := gin.Default()

	// 设置静态文件目录
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	// 路由设置
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "个人介绍",
		})
	})

	// 从环境变量获取端口，默认为 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 启动服务器，监听所有网络接口
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	log.Printf("服务器启动在 %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}
