package routes

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

var users []User

// 为方便测试，则直接通过init方式赋值。实际项目中，一般通过数据库等其它方式查询获取数据。
func init() {
	users = []User{
		{
			Username: "xcbeyond",
			Sex:      "男",
			Age:      18,
			Labels:   []string{"年轻", "帅"},
		},
		{
			Username: "niki",
			Sex:      "女",
			Age:      16,
			Labels:   []string{"漂亮"},
		},
	}
}

// SetupRouter 配置路由
func SetupRouter() *gin.Engine {
	route := gin.Default()

	route.GET("/user/:name", querUserHandler)
	// 其它更多路由

	return route
}

// handler
func querUserHandler(c *gin.Context) {
	name := c.Param("name")
	for _, user := range users {
		if user.Username == name {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusOK, "not found user :"+name)
}
