package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/dorieq/notification-center/internal/dispatcher"
	"github.com/segmentio/kafka-go"
)

func StartKafkaConsumer(ctx context.Context, topic, broker string) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:29092"},
		Topic:   topic,
		GroupID: "notify-group",
	})

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			log.Println("Kafka read error:", err)
			continue
		}
		var notif dispatcher.Notification
		if err := json.Unmarshal(m.Value, &notif); err != nil {
			log.Println("Invalid Kafka message:", err)
			continue
		}
		dispatcher.HandleNotification(notif)
	}
}
