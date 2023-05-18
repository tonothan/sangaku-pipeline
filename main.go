package main

import (
	"tonothan/sangaku-pipeline-server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Run routes
	routes.ManifestRoutes(router)

	router.Run(":8080")
}
