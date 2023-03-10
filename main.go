package main

import (
	"edu-management-system/router"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load("./config/db.yaml", "./config/minio.yaml")
	if err != nil {
		log.Fatal(err)
	}
	//db.MongoDBInit()
	router.Server()
}
