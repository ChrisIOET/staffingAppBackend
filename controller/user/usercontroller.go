package controller

import (
	"log"
	"net/http"
	"projectStaff/database"
	"projectStaff/model"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) { //cntexto datos q me manda la ruta.

	var usersList []model.User

	// GORM
	result := database.DB.Find(&usersList)

	if result.Error != nil {
		log.Printf("Error getting all categoories, reason: %v\n", result.Error)

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "something went wrong getting all categories",
		})
		return
	}
	// GIN -> preparo el json a devolver.
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All results",
		"data":    usersList,
	})
	return
}













