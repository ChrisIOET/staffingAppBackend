package routes

import (
	"github.com/gin-gonic/gin"
)

func CreateRoutes(router *gin.Engine)  {
	CategoryRoutes(router)
	UserRoutes(router)
}


