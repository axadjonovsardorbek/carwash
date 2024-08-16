package main

import (
	"gateway/api"
	"gateway/api/handler"
	cf "gateway/config"
	"gateway/grpc/clients"
	"gateway/kafka"
	"log"
)

func main() {
	config := cf.Load()

	services, err := clients.NewGrpcClients(&config)
	if err != nil {
		log.Fatalf("error while connecting clients. err: %s", err.Error())
	}

	broker := []string{config.KAFKA_HOST + config.KAFKA_PORT}
	kafka, err := kafka.NewKafkaProducer(broker)
	if err != nil {
		log.Fatalln("Failed to connect to Kafka", err)
		return
	}
	defer kafka.Close()

	engine := api.NewRouter(handler.NewHandler(services, &kafka))

	err = engine.Run(config.GATEWAY_PORT)
	if err != nil {
		log.Fatalf("error while running server. err: %s", err.Error())
	}
}
