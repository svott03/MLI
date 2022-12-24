package controllers

import (
	"MLI/responses"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"encoding/json"
	"os/exec"

	"github.com/gin-gonic/gin"
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
		// Retrieve File
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
		// Output to local file
		io.Copy(f, file)
		data, err2 := os.ReadFile("./files/" + handler.Filename)
		if err2 != nil {
			log.Fatal(err2)
		}
		// Send request to modelService
		req, _ := http.NewRequest("POST", "http://localhost:9090/uploadModel", bytes.NewBuffer(data))
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("client: error making http request: %s\n", err)
			return
		}
		fmt.Printf("client: status code: %d\n", res.StatusCode)
		// Read response from modelService and pass back to html
		bytes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			return
		}
		var out responses.BasicResponse
		err = json.Unmarshal(bytes, &out)
		if err != nil {
			log.Fatal(err)
			return
		}
		c.JSON(http.StatusOK, out)
	}
}

func UploadData() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("uploadData Controller...")
		// Retrieve File
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
		// Output to local file
		io.Copy(f, file)

		fmt.Println("file Uploaded")
		// Insert new training data into database
		cmd := exec.Command("zsh", "-c", "mongoimport --uri $MONGO_KEY -d MyDatabase --collection train_data --type=csv --headerline --file ./files/"+handler.Filename)
		out, err3 := cmd.Output()
		if err3 != nil {
			fmt.Println("could not run command: ", err3)
			return
		}
		fmt.Println("Output: " + string(out))
		c.JSON(http.StatusOK, responses.BasicResponse{Output: "Data Uploaded Successfully"})
	}
}

func UploadPredict() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("uploadPredict Controller...")
		// Retrieve File
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
		// Output to local file
		io.Copy(f, file)

		// Send file to modelService
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
		// Read response from modelService and pass back to html
		bytes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			return
		}
		var out responses.BasicResponse
		err = json.Unmarshal(bytes, &out)
		if err != nil {
			log.Fatal(err)
			return
		}
		if (res.StatusCode == 200) {
			c.JSON(http.StatusOK, out)
		} else {
			c.String(http.StatusInternalServerError, out.Output)
		}
	}
}

func Train() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Train Controller...")
		// Send http train request to modelService
		req, _ := http.NewRequest("GET", "http://localhost:9090/train", nil)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("client: error making http request: %s\n", err)
			return
		}
		fmt.Printf("client: status code: %d\n", res.StatusCode)
		// Read response from modelService and pass back to html
		bytes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			return
		}
		var out responses.BasicResponse
		err = json.Unmarshal(bytes, &out)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("OUT " + out.Output)
		if (res.StatusCode == 200) {
			c.JSON(http.StatusOK, out)
		} else {
			c.String(http.StatusInternalServerError, out.Output)
		}
	}
}
