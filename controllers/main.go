package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/saiteja111997/chatGPT_service/pkg/server"
)

func main() {
	port := "8080"
	r := gin.Default()
	// r.Use(server.EnableCors)
	r.POST("/create", server.GenerateFinalResultV2)
	r.POST("/test_webhook", server.TestWebhook)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
