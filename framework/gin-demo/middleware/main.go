package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化gin对象
	route := gin.Default()

	// 中间件注册
	route.Use(costTimeMiddleware())

	// 设置一个get请求，其URL为/hello，并实现简单的响应
	route.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world!",
		})
	})

	// // 局部中间件注册
	// route.GET("/hello", costTimeMiddleware(), func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "hello world!",
	// 	})
	// })

	// 启动服务
	route.Run()
}

func costTimeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求前获取当前时间
		nowTime := time.Now()

		// 请求处理
		c.Next()

		// 请求处理完获取花费的时间
		costTime := time.Since(nowTime)

		requestURL := c.Request.URL.String()
		fmt.Printf("the request URL %s cost %v\n", requestURL, costTime)
	}
}
