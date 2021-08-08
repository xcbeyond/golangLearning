package main

import (
	"math/rand"
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
		{
			Username: "niki",
			Sex:      "女",
			Age:      16,
			Labels:   []string{"漂亮"},
		},
	}

	// api分组
	api := route.Group("/api")
	{
		// v1分组
		v1 := api.Group("/v1")
		{
			v1.GET("/user/:name", func(c *gin.Context) {
				name := c.Param("name")
				for _, user := range users {
					if user.Username == name {
						c.JSON(http.StatusOK, user)
						return
					}
				}
				c.JSON(http.StatusOK, "not found user :"+name)
			})
		}

		// v2分组
		v2 := api.Group("/v2")
		{
			v2.GET("/user/:name", func(c *gin.Context) {
				name := c.Param("name")
				for _, user := range users {
					if user.Username == name {
						c.JSON(http.StatusOK, user)
						return
					}
				}

				// 如果没有查到，则随机返回一个
				user := users[rand.Intn(len(users)-1)]
				c.JSON(http.StatusOK, "not found user ["+name+"],but user ["+user.Username+"] is exist!")
			})
		}
	}

	route.Run()
}
