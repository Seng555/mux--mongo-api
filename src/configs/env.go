// configs/env.go
package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}
}

func EnvMongoURI() string {
	loadEnv()
	return os.Getenv("MONGOURI")
}

func EnvMongoDB() string {
	loadEnv()
	return os.Getenv("MONGODB")
}
func EnvPrivateKey() string {
	loadEnv()
	return os.Getenv("SERVER_KEY")
}
