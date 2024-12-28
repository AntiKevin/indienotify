package main

import (
	"log"

	"github.com/antikevin/indienotify/packages/api"
	"github.com/spf13/viper"
)

func main() {
	// Configurações
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// inicializa o servidor
	api.StartServer()
}
