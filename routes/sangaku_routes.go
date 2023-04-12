package routes

import (
	"pipeline/controllers"

	"github.com/gin-gonic/gin"
)

func SangakuRoutes(router *gin.Engine) {
	router.POST("/sangaku", controllers.CreateSangaku())
	router.POST("/images", controllers.UploadImages())
}
