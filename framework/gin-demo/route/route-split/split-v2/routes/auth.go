package routes

import (
	"github.com/gin-gonic/gin"
)

// AuthRegister route register
func AuthRegister(e *gin.Engine) {
	e.GET("/auth/login", loginHandler)
	e.POST("/auth/logout", logoutUserHanler)
	// ……
}

// loginHandler
func loginHandler(c *gin.Context) {

}

// logoutUserHanler
func logoutUserHanler(c *gin.Context) {

}
