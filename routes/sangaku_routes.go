package routes

import (
	"pipeline/controllers"

	"github.com/gin-gonic/gin"
)

func SangakuRoutes(router *gin.Engine) {
	router.POST("/sangaku", controllers.CreateSangaku())
	router.POST("/images", controllers.UploadImages())
	router.GET("/create-manifest", controllers.UploadImages())
	router.GET("/", controllers.Ping())
}
