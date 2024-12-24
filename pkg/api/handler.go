package api

import (
	"github.com/antikevin/indienotify/pkg/notifier"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Função para configurar as rotas da API
func SetupRoutes(r *gin.Engine) {
	r.POST("/api/send", SendNotificationHandler)
}

// Controlador para o endpoint /send
func SendNotificationHandler(c *gin.Context) {
	var request struct {
		Channel string `json:"channel" binding:"required"`
		Message string `json:"message" binding:"required"`
	}

	notifierService := notifier.NewNotifier(
		viper.GetString("redis.host"),
		viper.GetString("redis.port"),
		viper.GetString("redis.password"),
	)
	defer notifierService.Close()

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := notifierService.SendNotification(request.Channel, request.Message)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to send notification"})
		return
	}

	c.JSON(200, gin.H{"status": "Notification sent"})
}
