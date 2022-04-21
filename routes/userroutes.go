package routes

import (
	// "net/http"
	"projectStaff/controller/user"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/api/user", controller.GetUsers)
	router.GET("/api/user/:userId")
	router.PUT("/api/user/:userId")
	router.POST("/api/user")
	router.DELETE("/api/user/:userId")
}
