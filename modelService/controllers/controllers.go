package controllers

import (
	// "modelService/configs"
	// "modelService/models"
	"modelService/responses"
	"fmt"
	"net/http"
	// "time"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
)

func UploadModel() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("In modelService UploadModel")
		buf, err := io.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println("bad request")
      return
    }
		fmt.Println("Finished Read")
		err4 := os.WriteFile("./files/model.csv", buf, 0644)
		if err4 != nil {
			log.Fatal(err4)
			return
		}
		c.JSON(http.StatusOK, responses.BasicResponse{Output: "Model Uploaded Successfully"})
	}
}

func UploadPredict() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("In modelService UploadPredict")
		buf, err := io.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println("bad request")
      return
    }
		fmt.Println("Finished Read")
		err4 := os.WriteFile("./files/predict.csv", buf, 0644)
		if err4 != nil {
			log.Fatal(err4)
			return
		}
		//exec source code on predict input
		// cmd := os.exec.Command("mongoimport", "--uri $MONGO_KEY -d MyDatabase --collection meal_info --type=csv --headerline --file ~/Desktop/meal_info_new.csv")

		// err2 := cmd.Run()

		// if err2 != nil {
		// 	log.Fatal(err)
		// 	return
		// }

		// c.JSON(http.StatusOK, responses.BasicResponse{Output: "Data Uploaded"})
		c.JSON(http.StatusOK, responses.BasicResponse{Output: "Prediction Read, 85% accuracy"})
	}
}

func Train() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("In modelService Train")
		//check if source code exists

		// exec train model source code

		// send back train statistics
		c.JSON(http.StatusOK, responses.BasicResponse{Output: "Training Complete"})
	}
}
