package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// modulo de websocket para a API
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Função para atualizar a conexão HTTP para uma conexão websocket
func upgradeConnection(c *gin.Context) (*websocket.Conn, error) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return nil, err
	}
	return ws, nil
}
