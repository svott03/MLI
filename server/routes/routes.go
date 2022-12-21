package routes

import (
	"MLI/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", controllers.CreateUser())
	router.GET("/users", controllers.GetAllUsers())
	router.GET("/", controllers.GetMainPage())
	router.GET("/train", controllers.Train())
	router.POST("/uploadModel", controllers.UploadModel())
	router.POST("/uploadData", controllers.UploadData())
	router.POST("/uploadPredict", controllers.UploadPredict())
}