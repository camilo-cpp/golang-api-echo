package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/camilo-cpp/golang-api-echo/internal/server"
	"github.com/camilo-cpp/golang-api-echo/internal/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	uploadDataService := &services.UploadDataClient{}

	if err := uploadDataService.UploadData(); err != nil {
		log.Fatalf("Error uploading data 💀: %v", err)
	}

	if err := server.Start(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
