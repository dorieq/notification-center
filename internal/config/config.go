package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	KafkaBroker string
	KafkaTopic  string
	KafkaGroup  string
	WSOrigin    string
	JWTSecret   string
	LogLevel    string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	return &Config{
		Port:        getEnv("PORT", "8080"),
		KafkaBroker: getEnv("KAFKA_BROKER", "localhost:9092"),
		KafkaTopic:  getEnv("KAFKA_TOPIC", "notifications"),
		KafkaGroup:  getEnv("KAFKA_GROUP_ID", "notify-group"),
		WSOrigin:    getEnv("WS_ORIGIN", "*"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
