package controllers

import (
	"modelService/configs"
	// "modelService/models"
	"modelService/responses"
	// "context"
	"fmt"
	"net/http"
	// "time"
	"os"
	"os/exec"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

// func CreateUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		var user models.User
// 		defer cancel()

// 		//validate the request body
// 		if err := c.BindJSON(&user); err != nil {
// 			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 			return
// 		}

// 		//use the validator library to validate required fields
// 		if validationErr := validate.Struct(&user); validationErr != nil {
// 			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
// 			return
// 		}

// 		newUser := models.User{
// 			Id:       primitive.NewObjectID(),
// 			Name:     user.Name,
// 			Location: user.Location,
// 			Title:    user.Title,
// 		}

// 		result, err := userCollection.InsertOne(ctx, newUser)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 			return
// 		}

// 		c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
// 	}
// }

// func GetAllUsers() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		var users []models.User
// 		defer cancel()

// 		results, err := userCollection.Find(ctx, bson.M{})

// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 			return
// 		}

// 		//reading from the db in an optimal way
// 		defer results.Close(ctx)
// 		for results.Next(ctx) {
// 			var singleUser models.User
// 			if err = results.Decode(&singleUser); err != nil {
// 				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 			}

// 			users = append(users, singleUser)
// 		}

// 		c.JSON(http.StatusOK,
// 			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": users}},
// 		)
// 	}
// }


func UploadModel() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("uploadModel Controller...")
		c.Request.ParseMultipartForm(32 << 20)
		file, handler, err := c.Request.FormFile("file")
		if err != nil {
			fmt.Println("Err " + err.Error())
			return
		}
		fmt.Println("file Uploaded")
		defer file.Close()
		f, err := os.OpenFile("../files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		c.JSON(http.StatusOK, responses.BasicResponse{Output: "complete"})
	}
}

func UploadData() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("uploadData Controller...")
		c.Request.ParseMultipartForm(32 << 20)
		file, handler, err := c.Request.FormFile("file")
		if err != nil {
			fmt.Println("Err " + err.Error())
			return
		}
		fmt.Println("file Uploaded")
		defer file.Close()
		f, err := os.OpenFile("../files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)

		// put data into database
		cmd := exec.Command("mongoimport", "--uri $MONGO_KEY -d MyDatabase --collection meal_info --type=csv --headerline --file ~/Desktop/meal_info_new.csv")

    err2 := cmd.Run()

    if err2 != nil {
        log.Fatal(err)
    }


		c.JSON(http.StatusOK, responses.BasicResponse{Output: "Data Uploaded"})
	}
}

func UploadPredict() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("uploadPredict Controller...")
		c.Request.ParseMultipartForm(32 << 20)
		file, handler, err := c.Request.FormFile("file")
		if err != nil {
			fmt.Println("Err " + err.Error())
			return
		}
		fmt.Println("file Uploaded")
		defer file.Close()
		f, err := os.OpenFile("../files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		c.JSON(http.StatusOK, responses.BasicResponse{Output: "complete"})
	}
}

func Train() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Train Controller...")
		c.JSON(http.StatusOK, responses.BasicResponse{Output: "complete"})
	}
}
