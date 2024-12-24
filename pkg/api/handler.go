package api

import (
	"log"

	"github.com/antikevin/indienotify/pkg/notifier"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
)

// Função para configurar as rotas da API
func SetupRoutes(r *gin.Engine) {
	r.POST("/api/send", SendNotificationHandler)
	r.GET("/api/ws/receive/:channel", ReceiveNotificationHandler)
}

// Controlador para o endpoint de envio de notificações
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

// controlador para o endpoint de recebimento de notificações websocket
func ReceiveNotificationHandler(c *gin.Context) {

	ws, err := upgradeConnection(c)
	if err != nil {
		log.Printf("Failed to upgrade websocket connection: %v", err)
		c.JSON(400, gin.H{"error": "Failed to upgrade connection"})
		return
	}

	notifierService := notifier.NewNotifier(
		viper.GetString("redis.host"),
		viper.GetString("redis.port"),
		viper.GetString("redis.password"),
	)
	defer notifierService.Close()

	//obtém o canal pelo parametro da URL
	channel := c.Params.ByName("channel")
	if !notifierService.IsValidChannel(channel) {
		c.JSON(400, gin.H{"error": "Channel not specified"})
		return
	}

	//adiciona a conexão ao registro
	connections := make(map[*websocket.Conn]bool)
	connections[ws] = true

	//escuta por mensagens
	for {
		//recebe a mensagem
		msg, err := notifierService.ReadNotifications(channel)

		//se houver um erro, fecha a conexão
		if err != nil {
			ws.WriteMessage(websocket.TextMessage, []byte("failed to connect to notifications reader"))
			delete(connections, ws)
			break
		}

		// mostra a mensagem recebida
		if err := ws.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			log.Printf("Failed to write message: %v", err)
			delete(connections, ws)
			break
		}
	}
}
