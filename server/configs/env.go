package configs

import (
	// "fmt"
	// "os"
)

func EnvMongoURI() string {
	// err := godotenv.Load()
  // if err != nil {
  //   log.Fatal("Error loading .env file")
  // }
	// token := os.Getenv("MONGOURI")
	// fmt.Println(token)
	// return token
	return "mongodb+srv://svott:Burrito77@cluster0.v1wrvyg.mongodb.net/?retryWrites=true&w=majority"
}