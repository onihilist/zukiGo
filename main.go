package main

import (
	"log"
	"os"
	"zukiapi/handlers"
	"zukiapi/services"

	"github.com/gin-gonic/gin"
)

func main() {
	apiKey := os.Getenv("ZUKI_API_KEY")
	if apiKey == "" {
		log.Fatal("[ERROR] - ZUKI_API_KEY must be set in environment variables")
	}

	services.InitZuki(apiKey)

	r := gin.Default()
	r.POST("/chat", handlers.ChatHandler)
	r.POST("/image", handlers.ImageHandler)
	r.Run(":8080")
}
