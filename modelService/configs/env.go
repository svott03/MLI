package configs

import (
	"os"
)

func EnvMongoURI() string {
	token := os.Getenv("MONGO_KEY")
	return token
}

func EnvMongoDb() string {
	token := os.Getenv("MongoDB")
	return token
}

func EnvMongoTrainCollection() string {
	token := os.Getenv("MongoTrainData")
	return token
}