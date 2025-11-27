package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func MyEnv(str string) interface{} {
	InitEnv()
	return os.Getenv(str)
}
