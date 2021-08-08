package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	// 设置一个get请求，其URL为/hello，并实现简单的响应
	route.GET("/get", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "this is a get method response!",
		})
	})

	// 具体实现可单独定义一个函数
	route.POST("/post", postHandler)

	route.PUT("/put", func(c *gin.Context) {

	})

	route.PATCH("/patch", func(c *gin.Context) {

	})

	route.DELETE("/delete", func(c *gin.Context) {

	})

	// ……

	route.Run()
}

func postHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "this is a post method response!")
}
