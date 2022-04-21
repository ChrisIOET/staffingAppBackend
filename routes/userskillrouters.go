package routes

import (
	// "net/http"
	"github.com/gin-gonic/gin"
)

func UserSkillRoutes(router *gin.Engine){
	router.GET("/api/userskill")
	router.GET("/api/userskill/:userskillId")
	router.PUT("/api/userskill/:userskillId")
	router.POST("/api/userskill")
	router.DELETE("/api/userskill/:userskillId")
}






