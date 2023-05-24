package routes

import (
	"tonothan/sangaku-pipeline-server/controllers"

	"github.com/gin-gonic/gin"
)

func ManifestRoutes(router *gin.Engine) {
	router.GET("/", controllers.Ping())
	router.GET("/manifest/:manifestId", controllers.GetManifestMetadata())
	router.POST("/manifest", controllers.CreateManifestMetadata())
}
