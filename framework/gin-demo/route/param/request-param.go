package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string   `json:"username"`
	Sex      string   `json:"sex"`
	Age      int      `json:"age"`
	Labels   []string `json:"lalels"`
}

func main() {
	route := gin.Default()

	users := []User{
		{
			Username: "xcbeyond",
			Sex:      "男",
			Age:      18,
			Labels:   []string{"年轻", "帅"},
		},
	}

	// 请求path中存在参数
	route.GET("/user/:name", func(c *gin.Context) {
		// 获取请求path中的参数
		name := c.Param("name")
		for _, user := range users {
			if user.Username == name {
				c.JSON(http.StatusOK, user)
				return
			}
		}
		c.JSON(http.StatusOK, "not found user "+name)
	})

	// 查询参数，如：/user?name=xcbeyond
	route.GET("/user", func(c *gin.Context) {
		// 获取中的参数
		name := c.Query("name")
		for _, user := range users {
			if user.Username == name {
				c.JSON(http.StatusOK, user)
				return
			}
		}
		c.JSON(http.StatusOK, "not found user "+name)
	})

	route.Run()
}
