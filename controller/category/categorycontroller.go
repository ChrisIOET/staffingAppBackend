package controller

import (
	"log"
	"net/http"
	"projectStaff/database"
	"projectStaff/model"
	"time"
	// "fmt"
	"github.com/gin-gonic/gin"
	// "github.com/go-pg/pg/v9"
)

func GetCategories(c *gin.Context) { //cntexto datos q me manda la ruta.

	var categoriesList []model.Category

	// GORM
	result := database.DB.Find(&categoriesList)

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
		"data":    categoriesList,
	})
	return
}

func CreateCategory(c *gin.Context) {

	var category model.Category

	//obtener json y el id generado
	c.BindJSON(&category)

	name  := category.Name //el unico dato q me mandan en el api.
	CreatedAt := time.Now()
	UpdatedAt := time.Now()

	newCategory := model.Category{Name: name, CreatedAt: CreatedAt, UpdatedAt: UpdatedAt}

	result := database.DB.Select("name", "CreatedAt", "UpdatedAt").Create(&newCategory)

	if result.Error != nil {
		log.Printf("Error while inserting new cateogry into db, Reason: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Category created Successfully",
	})
	return
}

func DeleteCategories(c *gin.Context) {

	categoryId := c.Param("categoryId")

	result := database.DB.Unscoped().Delete(&model.Category{}, categoryId) //vamos a borrar la llave primaria cateogry ID

	if result.Error != nil {
		log.Printf("Error while deleting a single category, Reason: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "category deleted successfully",
	})
	return
}


func EditCategories(c *gin.Context) {

	categoryId := c.Param("categoryId")

	//las 2 lineas de abajo son necesarias
	var category model.Category
	c.BindJSON(&category)	// ---> super importante para que me cambie el value dentro de la DB
	Name := category.Name

	result := database.DB.Find(&model.Category{}).Where("id = ?",categoryId).Select("name").Update("name", Name)

	if result.Error != nil {
		log.Printf("Error while deleting a single edit, Reason: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Category Edited Successfully",
	})
	return
}


