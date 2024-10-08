package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	GATEWAY_PORT string
	BOOKING_PORT string

	BOOKING_HOST string

	KAFKA_HOST string
	KAFKA_PORT string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.GATEWAY_PORT = cast.ToString(coalesce("GATEWAY_PORT", ":8081"))
	config.BOOKING_PORT = cast.ToString(coalesce("BOOKING_PORT", ":50051"))
	config.BOOKING_HOST = cast.ToString(coalesce("BOOKING_HOST", "memory-service"))

	config.KAFKA_HOST = cast.ToString(coalesce("KAFKA_HOST", "localhost"))
	config.KAFKA_PORT = cast.ToString(coalesce("KAFKA_PORT", ":9092"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
