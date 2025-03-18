package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kalpesh-vala/go-chat-app/controllers"
)

func AuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}
}
