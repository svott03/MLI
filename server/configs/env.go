package configs

import (
	"fmt"
	"os"
)

func EnvMongoURI() string {
	token := os.Getenv("MONGO_KEY")
	fmt.Println(token)
	return token
}