package main

import (
	"modelService/routes"
	"modelService/configs"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	configs.ConnectDB()
	routes.UserRoute(router)
	router.Run("localhost:9090")
}