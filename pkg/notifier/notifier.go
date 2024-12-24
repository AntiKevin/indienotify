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

func (m *Notifier) IsValidChannel(channel string) bool {
	return channel != "" && channel != " " && channel != "\n" && channel != "\t"
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

func (n *Notifier) ReadNotifications(channel string) (string, error) {
	ctx := context.Background()
	pubsub := n.redisClient.Subscribe(ctx, channel)
	defer pubsub.Close()

	msg, err := pubsub.ReceiveMessage(ctx)
	return msg.Payload, err
}

func (n *Notifier) Close() error {
	return n.redisClient.Close()
}
