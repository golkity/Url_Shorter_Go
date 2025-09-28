package kafka

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
	log    *slog.Logger
}

type UserCreatedEvent struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func NewProducer(brokers []string, topic string, log *slog.Logger) *Producer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	return &Producer{writer: writer, log: log}
}

func (p *Producer) PublishUserCreated(ctx context.Context, userID, email string) {
	const op = "kafka.Producer.PublishUserCreated"

	event := UserCreatedEvent{ID: userID, Email: email}
	eventBytes, err := json.Marshal(event)
	if err != nil {
		p.log.Error("failed to marshal event", "op", op, "error", err)
		return
	}

	err = p.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(userID),
		Value: eventBytes,
	})

	if err != nil {
		p.log.Error("failed to write message to kafka", "op", op, "error", err)
	} else {
		p.log.Info("sent 'user.created' event to kafka", "user_id", userID)
	}
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
