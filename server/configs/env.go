package configs

import (
	"os"
)

func EnvMongoURI() string {
	token := os.Getenv("MONGO_KEY")
	return token
}