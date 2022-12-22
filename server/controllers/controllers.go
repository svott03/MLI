package controllers

import (
	// "MLI/configs"
	// "MLI/models"
	"MLI/responses"
	"bytes"
	// "context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	// "time"
	"os/exec"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
)

func GetMainPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("In main controller...")
		c.HTML(http.StatusOK, "index.html", nil)
	}
}

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
		f, err := os.OpenFile("./files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		data, err2 := os.ReadFile("./files/" + handler.Filename)
		if err2 != nil {
			log.Fatal(err2)
		}
		req, _ := http.NewRequest("POST", "http://localhost:9090/uploadPredict", bytes.NewBuffer(data))
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("client: error making http request: %s\n", err)
			return
		}
		fmt.Printf("client: status code: %d\n", res.StatusCode)
		c.JSON(http.StatusOK, res.Body)
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
		defer file.Close()
		f, err := os.OpenFile("./files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)

		fmt.Println("file Uploaded")
		// put data into database
		
		// cmd := exec.Command("mongoimport --uri $MONGO_KEY -d MyDatabase --collection meal_info --type=csv --headerline --file ./files/" + handler.Filename)
		cmd := exec.Command("zsh", "-c", "mongoimport --uri $MONGO_KEY -d MyDatabase --collection meal_info --type=csv --headerline --file ./files/" + handler.Filename)
		out, err3 := cmd.Output()
		if err3 != nil {
			// if there was any error, print it here
			fmt.Println("could not run command: ", err3)
		}
		fmt.Println("Output: ", string(out))
		// err2 := cmd.Run()
		// if err2 != nil {
		// 	log.Fatal(err2)
		// 	return
		// }

		c.JSON(http.StatusOK, responses.BasicResponse{Output: "Data Uploaded Successfully"})
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
		f, err := os.OpenFile("./files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)

		//send file to modelService
		data, err2 := os.ReadFile("./files/" + handler.Filename)
		if err2 != nil {
			log.Fatal(err2)
		}
		req, _ := http.NewRequest("POST", "http://localhost:9090/uploadPredict", bytes.NewBuffer(data))
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("client: error making http request: %s\n", err)
			return
		}
		fmt.Printf("client: status code: %d\n", res.StatusCode)
		c.JSON(http.StatusOK, responses.BasicResponse{Output: "complete"})
	}
}

func Train() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Train Controller...")
		//send http train request to modelService
		c.JSON(http.StatusOK, responses.BasicResponse{Output: "complete"})
	}
}
