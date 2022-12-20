package main

import (
	"MLI/configs"
	"MLI/routes"
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	router := gin.Default()

	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"data": "Hello from Gin-gonic & mongoDB",
	// 	})
	// })
	router.LoadHTMLGlob("templates/*.html")
	fmt.Println("In Main")
	configs.ConnectDB()
	routes.UserRoute(router)
	router.Run("localhost:8080")
}