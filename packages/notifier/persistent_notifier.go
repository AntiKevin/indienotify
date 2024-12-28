package notifier

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// Notification é a estrutura de dados para uma notificação
type Notification struct {
	ID        string `json:"id"`
	Topic     string `json:"topic"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

// PersistentNotifier é a interface para um serviço de notificação
type PersistentNotifier struct {
	RedisClient *redis.Client
	StreamName  string
}

// NewPersistentNotifier cria um novo serviço de notificação persistente
func NewPersistentNotifier(client *redis.Client, streamName string) *PersistentNotifier {
	return &PersistentNotifier{
		RedisClient: client,
		StreamName:  streamName,
	}
}

// publica uma notificação persistente
func (pn *PersistentNotifier) Publish(ctx context.Context, notification Notification) error {
	_, err := pn.RedisClient.XAdd(ctx, &redis.XAddArgs{
		Stream: pn.StreamName,
		Values: map[string]interface{}{
			"id":        notification.ID,
			"topic":     notification.Topic,
			"message":   notification.Message,
			"timestamp": notification.Timestamp,
		},
	}).Result()
	return err
}

// consome notificações persistentes e retorna uma lista de mensagens
func (pn *PersistentNotifier) Consume(ctx context.Context, lastID string) ([]redis.XMessage, error) {
	result, err := pn.RedisClient.XRead(ctx, &redis.XReadArgs{
		Streams: []string{pn.StreamName, lastID},
		Count:   10,
		Block:   0,
	}).Result()
	if err != nil {
		return nil, err
	}
	return result[0].Messages, nil
}
