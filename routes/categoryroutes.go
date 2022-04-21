package routes

import (
	// "net/http"
	"projectStaff/controller/category"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine) {
	router.GET("/api/category", controller.GetCategories)
	// router.GET("/api/category/:categoryId")
	router.PUT("/api/category/:categoryId", controller.EditCategories)
	router.POST("/api/category", controller.CreateCategory)
	router.DELETE("/api/category/:categoryId", controller.DeleteCategories)
}
