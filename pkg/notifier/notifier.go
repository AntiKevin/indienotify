package notifier

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type Notifier struct {
	redisClient *redis.Client
}

func NewNotifier(redisHost, redisPort, redisPassword string) *Notifier {
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword,
		DB:       0,
	})

	return &Notifier{redisClient: client}
}

func (n *Notifier) SendNotification(channel, message string) error {
	ctx := context.Background()
	if err := n.redisClient.Publish(ctx, channel, message).Err(); err != nil {
		log.Printf("Erro ao enviar notificação: %v", err)
		return err
	}
	log.Printf("Notificação enviada: %s -> %s", channel, message)
	return nil
}

func (n *Notifier) Close() error {
	return n.redisClient.Close()
}
