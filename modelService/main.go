package main

import (
	"modelService/routes"
	"modelService/configs"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"data": "Hello from Gin-gonic & mongoDB",
	// 	})
	// })
	configs.ConnectDB()
	routes.UserRoute(router)
	router.Run("localhost:9090")
}