package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kalpesh-vala/go-chat-app/config"
	"github.com/kalpesh-vala/go-chat-app/controllers"
	"github.com/kalpesh-vala/go-chat-app/routes"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	// Initialize MongoDB (Ensure it's connected)
	if config.DB == nil {
		log.Fatal("Database is not initialized. Exiting...")
	}

	// Initialize Redis
	config.InitRedis()

	// Start Redis Subscription in a separate goroutine to listen for incoming messages
	go controllers.SubscribeToMessages()

	// Initialize Gin Router
	r := gin.Default()

	// Set trusted proxy (Avoiding 'nil' to prevent errors)
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatalf("Error setting trusted proxies: %v", err)
	}

	// Load authentication routes (✅ Keeping it as per your requirement)
	routes.AuthRoutes(r)

	// Load chat routes (✅ Keeping chat functionality)
	routes.ChatRoutes(r)

	fmt.Println("✅ Server is running on http://localhost:8080")
	r.Run(":8080") // Start server
}
