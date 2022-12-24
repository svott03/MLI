package controllers

import (
	"modelService/configs"
	"modelService/models"
	"fmt"
	"modelService/responses"
	"net/http"
	"time"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection = configs.GetCollection(configs.DB, "train_data")

func UploadModel() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("In modelService UploadModel")
		buf, err := io.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println("bad request")
			return
		}
		fmt.Println("Finished Read")
		err4 := os.WriteFile("./files/model.py", buf, 0644)
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
		//check if source code exists
		if _, err := os.Stat("./files/model.py"); err == nil {
			fmt.Printf("File exists\n")
		} else {
			fmt.Printf("File does not exist\n")
			c.JSON(http.StatusConflict, responses.BasicResponse{Output: "Model Source Code needs to be imported."})
			return
		}
		// exec train model source code
		cmd := exec.Command("zsh", "-c", "python3 ./files/prediction.py")
		out, err3 := cmd.Output()
		if err3 != nil {
			fmt.Println("could not run command: ", err3)
			return
		}
		c.JSON(http.StatusOK, responses.BasicResponse{Output: "Prediction: " + string(out)})
	}
}

func Train() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("In modelService Train")
		//check if source code exists
		if _, err := os.Stat("./files/model.py"); err == nil {
			fmt.Printf("File exists\n")
		} else {
			fmt.Printf("File does not exist\n")
			c.JSON(http.StatusConflict, responses.BasicResponse{Output: "Model Source Code needs to be imported."})
			return
		}
		//grab data
		res, _ := os.ReadFile("./files/numRecords.txt")
		prevRecords, _ := strconv.Atoi(string(res))
		fmt.Println(prevRecords)
		// collection.find().skip(collection.count() - prevRecords)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var train_data []models.Instance
		defer cancel()

		results, err := collection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BasicResponse{Output: "Error in Loading Training Data"})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var instance models.Instance
			var m bson.M
			err = results.Decode(&m)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.BasicResponse{Output: "Error in Model Source Code"})
				return
			}
			fmt.Println(m)
			bsonBytes, _ := bson.Marshal(m)
			_ = bson.Unmarshal(bsonBytes, &instance)
			fmt.Println(instance)
			train_data = append(train_data, instance)
		}

		// exec train model source code
		cmd := exec.Command("zsh", "-c", "python3 ./files/model.py")
		out, err3 := cmd.Output()
		if err3 != nil {
			fmt.Println("could not run command: ", err3)
			c.JSON(http.StatusInternalServerError, responses.BasicResponse{Output: "Error in Model Source Code"})
			return
		}
		fmt.Println("Output: ", string(out))
		// send back train statistics
		c.JSON(http.StatusOK, responses.BasicResponse{Output: "Training Complete. Output: " + string(out)})
	}
}
