package main

import (
	"fmt"
	"log"
	// "projectStaff/controller"
	"projectStaff/database"
	// "projectStaff/repository"
	"net/http"
	"os"
	"projectStaff/model"

	"projectStaff/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8081"
	}

	database.DB = database.GetDatabase()

	router := gin.Default()
	// Automigracion

	database.DB.AutoMigrate(&model.Skill{})
	database.DB.AutoMigrate(&model.User{})
	database.DB.AutoMigrate(&model.UserSkill{})
	database.DB.AutoMigrate(&model.Category{})

	// rutas

	routes.CategoryRoutes(router)
	routes.UserRoutes(router)
	log.Fatal(router.Run(":8081"))

	// db := database.GetDatabase()
	// repos := repository.InitRepositories(db)
	// controllers := controller.InitControllers(repos)

	fmt.Println("server is started at: http://localhost:/" + port + "/")
	http.ListenAndServe(":"+port, nil)
}
