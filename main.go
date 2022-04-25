package main

import (
	"log"
	"projectStaff/database"
	"projectStaff/internal/envvar"
	"projectStaff/routes"
	"github.com/gin-gonic/gin"
)


func main() {
	PORT := envvar.ApiPort()

	database.DB = database.GetDatabase()
	database.MigrateDatabaseTables()

	router := gin.Default()
	routes.CreateRoutes(router)

	log.Fatal(router.Run(":"+PORT))
}
