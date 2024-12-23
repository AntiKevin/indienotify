package main

import (
	"log"

	"github.com/antikevin/indienotify/pkg/notifier"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// Configurações
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Erro ao carregar config.yaml: %v", err)
	}

	notifierService := notifier.NewNotifier(
		viper.GetString("redis.host"),
		viper.GetString("redis.port"),
		viper.GetString("redis.password"),
	)

	// Inicializa o servidor
	r := gin.Default()

	r.POST("/send", func(c *gin.Context) {
		var request struct {
			Channel string `json:"channel" binding:"required"`
			Message string `json:"message" binding:"required"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := notifierService.SendNotification(request.Channel, request.Message); err != nil {
			c.JSON(500, gin.H{"error": "Failed to send notification"})
			return
		}

		c.JSON(200, gin.H{"status": "notification sent"})
	})

	// Exemplo de rota básica
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Inicia o servidor na porta configurada
	port := viper.GetString("server.port")
	if port == "" {
		port = "8080" // Porta padrão
	}

	log.Printf("IndieNotify rodando na porta %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
