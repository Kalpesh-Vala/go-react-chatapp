package services

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/kalpesh-vala/go-chat-app/config"
	"github.com/kalpesh-vala/go-chat-app/models"
)

// SaveMessage stores chat messages in MongoDB
func SaveMessage(chat models.Chat) error {
	collection := config.GetCollection("messages")
	chat.Timestamp = time.Now()
	_, err := collection.InsertOne(context.Background(), chat)
	return err
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allows connections from any origin
	},
}

// HandleWebSocket processes WebSocket messages
func HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}

		// Convert to JSON before publishing
		var chatMessage models.Chat
		err = json.Unmarshal(msg, &chatMessage)
		if err != nil {
			log.Println("Invalid JSON format:", err)
			continue
		}

		// Ensure required fields are present
		if chatMessage.From == "" || chatMessage.To == "" || chatMessage.Msg == "" {
			log.Println("Invalid chat message: Missing required fields")
			continue
		}

		chatMessage.Timestamp = time.Now()

		messageJSON, err := json.Marshal(chatMessage)
		if err != nil {
			log.Println("Error marshalling message:", err)
			continue
		}

		// Publish message to Redis
		err = config.PublishMessage("chat", string(messageJSON))
		if err != nil {
			log.Println("Error publishing message:", err)
		}
	}
}
