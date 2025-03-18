package controllers

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kalpesh-vala/go-chat-app/config"
	"github.com/kalpesh-vala/go-chat-app/models"
	"github.com/kalpesh-vala/go-chat-app/services"
)

// HandleChat handles WebSocket requests
func HandleChat(c *gin.Context) {
	services.HandleWebSocket(c)
}

// SubscribeToMessages listens for messages from Redis
func SubscribeToMessages() {
	pubsub := config.SubscribeToChannel("chat")
	defer pubsub.Close()

	for msg := range pubsub.Channel() {
		var chatMessage models.Chat

		// Decode JSON message correctly
		err := json.Unmarshal([]byte(msg.Payload), &chatMessage)
		if err != nil {
			log.Println("Error parsing Redis message:", err)
			continue
		}

		// Store message in MongoDB
		err = services.SaveMessage(chatMessage)
		if err != nil {
			log.Println("Error saving message to MongoDB:", err)
		} else {
			log.Println("Message saved to MongoDB:", chatMessage)
		}
	}
}
