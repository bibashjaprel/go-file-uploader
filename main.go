package main

import (
	"log"
	"os"

	"github.com/bibashjaprel/go-file-uploader/localstorage"
	"github.com/bibashjaprel/go-file-uploader/s3storage"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if it exists (optional in production)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Ensure upload directory exists
	if err := os.MkdirAll(localstorage.UploadDir, os.ModePerm); err != nil {
		log.Fatalf("failed to create upload dir: %v", err)
	}
	router := gin.Default()

	// Max upload size: 10MB
	router.MaxMultipartMemory = 10 << 20 // 10 MiB

	router.POST("/upload", localstorage.HandleUpload)
	router.GET("/uploads/:filename", localstorage.HandleGetFile)
	router.POST("/s3/upload", s3storage.HandleUpload)

	log.Println("Server running at http://localhost:8080")
	router.Run(":8080")
}
