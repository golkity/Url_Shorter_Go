package kafka

import (
	"context"
	"encoding/json"
	"log/slog"

	"url-shortener/internal/repository"

	"github.com/segmentio/kafka-go"
)

type UserCreatedEvent struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func RunConsumer(ctx context.Context, log *slog.Logger, brokers []string, userRepo repository.UserRepository) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   "users.created",
		GroupID: "url-shortener-group",
	})

	log.Info("Starting Kafka consumer...")
	defer reader.Close()

	for {
		msg, err := reader.FetchMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				break
			}
			log.Error("could not fetch message", "error", err)
			continue
		}

		var event UserCreatedEvent
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			log.Error("could not unmarshal message", "error", err)
			reader.CommitMessages(ctx, msg)
			continue
		}

		log.Info("Received user.created event", "user_id", event.ID)

		if err := userRepo.CreateUser(ctx, event.ID); err != nil {
			log.Error("failed to process user.created event", "error", err)
			continue
		}

		if err := reader.CommitMessages(ctx, msg); err != nil {
			log.Error("could not commit message", "error", err)
		}
	}
	log.Info("Stopping Kafka consumer...")
}
