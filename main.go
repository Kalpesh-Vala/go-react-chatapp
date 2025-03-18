package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kalpesh-vala/go-chat-app/config"
	"github.com/kalpesh-vala/go-chat-app/routes"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	config.ConnectDB()

	if config.DB == nil {
		log.Fatal("Database is not initialized. Exiting...")
	}

	r := gin.Default()

	err := r.SetTrustedProxies([]string{"127.0.0.1"}) // Set a trusted IP instead of nil
	if err != nil {
		log.Fatalf("Error setting trusted proxies: %v", err)
	}

	routes.AuthRoutes(r)

	fmt.Println("Server is running on http://localhost:8080")
	r.Run(":8080")
}
