package main

import (
	"context"
	"log"
	"net/http"

	"github.com/dorieq/notification-center/internal/config"
	"github.com/dorieq/notification-center/internal/kafka"
	"github.com/dorieq/notification-center/internal/websocket"
)

func main() {
	cfg := config.LoadConfig()

	go websocket.HubInstance.Run()
	ctx := context.Background()

	go kafka.StartKafkaConsumer(ctx, cfg.KafkaTopic, cfg.KafkaBroker)

	http.HandleFunc("/ws", websocket.HandleWebSocket)

	log.Println("Server running on port", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
