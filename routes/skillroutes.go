package routes

import (
	"github.com/gin-gonic/gin"
)

func SkillRoutes(router *gin.Engine) {
	router.GET("/api/skill")
	router.GET("/api/skill/:skillId")
	router.PUT("/api/skill/:skillId")
	router.POST("/api/skill")
	router.DELETE("/api/skill/:skillId")
}
