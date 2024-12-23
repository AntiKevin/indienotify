package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func StartServer() {
	// Inicializa o servidor
	r := gin.Default()

	// Exemplo de rota básica
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Configura as rotas da API
	SetupRoutes(r)

	// Inicia o servidor na porta configurada
	port := viper.GetString("server.port")
	if port == "" {
		port = "8080" // Porta padrão
	}

	log.Printf("IndieNotify running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
