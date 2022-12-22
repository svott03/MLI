package routes

import (
	"modelService/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/train", controllers.Train())
	router.POST("/uploadModel", controllers.UploadModel())
	router.POST("/uploadPredict", controllers.UploadPredict())
}