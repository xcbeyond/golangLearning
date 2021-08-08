package routes

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter 配置路由
func SetupRouter() *gin.Engine {
	route := gin.Default()

	// other config

	// register all route.
	UserRegister(route)
	AuthRegister(route)
	// ……

	return route
}
