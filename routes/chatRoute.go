package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kalpesh-vala/go-chat-app/controllers"
)

func ChatRoutes(router *gin.Engine) {
	router.GET("/ws", controllers.HandleChat)
}
