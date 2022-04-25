package controller

import (
	"log"
	"net/http"
	"projectStaff/database"
	"projectStaff/model"
	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context) { 

	var usersList []model.User

	result := database.DB.Find(&usersList)

	if result.Error != nil {
		log.Printf("Error getting all categoories, reason: %v\n", result.Error)

		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "something went wrong getting all categories",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All results",
		"data":    usersList,
	})
	return
}
