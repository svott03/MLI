package controllers

import (
	"context"
	"fmt"
	"io"
	"log"
	"modelService/configs"
	"modelService/models"
	"modelService/responses"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection = configs.GetCollection(configs.DB, "train_data")

func UploadModel() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("In modelService UploadModel")
		buf, err := io.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println("bad request")
			c.JSON(http.StatusInternalServerError, responses.BasicResponse{Output: "Error in Uploading Model"})
			return
		}
		fmt.Println("Finished Read")
		err4 := os.WriteFile("./files/model.py", buf, 0644)
		if err4 != nil {
			log.Fatal(err4)
			c.JSON(http.StatusInternalServerError, responses.BasicResponse{Output: "Error in Writing File"})
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
			c.JSON(http.StatusInternalServerError, responses.BasicResponse{Output: "Error in Uploading Prediction"})
			return
		}
		fmt.Println("Finished Read")
		err4 := os.WriteFile("./files/predict.csv", buf, 0644)
		if err4 != nil {
			log.Fatal(err4)
			c.JSON(http.StatusInternalServerError, responses.BasicResponse{Output: "Error in Reading File"})
			return
		}
		//check if source code exists
		if _, err := os.Stat("./files/model.py"); err == nil {
			fmt.Printf("File exists\n")
		} else {
			fmt.Printf("File does not exist\n")
			c.JSON(http.StatusInternalServerError, responses.BasicResponse{Output: "Model Source Code needs to be imported."})
			return
		}
		// exec train model source code
		cmd := exec.Command("zsh", "-c", "python3 ./files/prediction.py")
		out, err3 := cmd.Output()
		if err3 != nil {
			fmt.Println("could not run command: ", err3)
			c.JSON(http.StatusInternalServerError, responses.BasicResponse{Output: "Error in Executing prediction.py"})
			return	
		}
		fmt.Println("Prediction: " + string(out))
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
			c.JSON(http.StatusInternalServerError, responses.BasicResponse{Output: "Model Source Code needs to be imported."})
			return
		}
		//grab data
		res, _ := os.ReadFile("./files/numRecords.txt")
		s := string(res)
		s = s[:len(s)-1]
		prevRecords, _ := strconv.Atoi(s)
		fmt.Println(prevRecords)
		// collection.find().skip(collection.count() - prevRecords)

		ctx, cancel := context.WithTimeout(context.Background(), 80*time.Second)
		var train_data []models.Instance
		defer cancel()
		// myOptions := options.Find()
		// myOptions.SetSort(bson.M{"$natural": prevRecords})
		opts := options.Find().SetSort(bson.D{{"$natural", 1}}).SetSkip(int64(prevRecords))
		results, err := collection.Find(ctx, bson.M{},
			opts,
		)
		// results.skip(prevRecords)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BasicResponse{Output: "Error in Loading Training Data"})
			return
		}

		//reading from the db in an optimal way
		w, _ := os.OpenFile("./files/train.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		newly_added := 0
		defer results.Close(ctx)
		for results.Next(ctx) {
			var instance models.Instance
			var m bson.M
			err = results.Decode(&m)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.BasicResponse{Output: "Error in Loading Training Data"})
				return
			}
			bsonBytes, _ := bson.Marshal(m)
			_ = bson.Unmarshal(bsonBytes, &instance)
			train_data = append(train_data, instance)
			newly_added++
			if prevRecords != 0 {
				// new_inst, _ := gocsv.MarshalString(&s)
				var s string = "\n"
				v := reflect.ValueOf(instance)

				for i := 0; i < v.NumField(); i++ {
					temp := fmt.Sprintf("%v,", v.Field(i).Interface())
					s += temp
				}
				s = s[:len(s)-1]
				fmt.Println(s)
				if _, err := w.WriteString(s); err != nil {
					//write failed do something
					log.Fatal("Error appending new data")
					c.JSON(http.StatusInternalServerError, responses.BasicResponse{Output: "Error in Writing to train.csv"})
					return
				}
			}
		}
		fmt.Println(newly_added)
		prevRecords += newly_added
		// Write size of db to file
		f, _ := os.Create("./files/numRecords.txt")
		defer f.Close()
		_, _ = f.WriteString(fmt.Sprintf("%d\n", prevRecords))

		fmt.Println(len(train_data))
		//output to csv file
		// to download file inside downloads folder
		file, err := os.OpenFile("./files/train.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, responses.BasicResponse{Output: "Error in opening train.csv"})
			return
		}
		defer file.Close()
		// remove header line
		if prevRecords-newly_added == 0 {
			gocsv.MarshalFile(&train_data, file)
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
