package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dorieq/notification-center/internal/config"
	"github.com/dorieq/notification-center/internal/kafka"
	"github.com/dorieq/notification-center/internal/websocket"
)

func main() {
	cfg := config.LoadConfig()

	go websocket.HubInstance.Run()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go kafka.StartKafkaConsumer(ctx, cfg.KafkaTopic, cfg.KafkaBroker)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: http.DefaultServeMux,
	}

	http.HandleFunc("/ws", websocket.HandleWebSocket)

	go func() {
		log.Println("Server running on port", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v\n", err)
		}
	}()
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutting down server...")

	shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 5*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("HTTP shutdown error: %v\n", err)
	}

	log.Println("Server exited properly")

}
