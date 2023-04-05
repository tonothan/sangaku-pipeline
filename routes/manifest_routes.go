package routes

import (
	"gin-mongo-api/controllers"
	"github.com/gin-gonic/gin"
)

func ManifestRoutes(router *gin.Engine) {
	router.POST("/manifest", controllers.CreateManifest())

	router.GET("/manifests", controllers.GetAllManifests())
}
