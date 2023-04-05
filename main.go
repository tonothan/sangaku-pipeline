package main

import (
	"gin-mongo-api/configs"
	"gin-mongo-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()

	routes.ManifestRoutes(router)

	router.Run("localhost:6000")
}
