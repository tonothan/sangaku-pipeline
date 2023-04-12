package main

import (
	"pipeline/configs"
	"pipeline/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()

	routes.SangakuRoutes(router)

	router.Run("localhost:6000")
}
