package main

import (
	"gateway/api"
	"gateway/api/handler"
	cf "gateway/config"
	"gateway/grpc/clients"
	"log"
)

func main() {
	config := cf.Load()

	services, err := clients.NewGrpcClients(&config)
	if err != nil {
		log.Fatalf("error while connecting clients. err: %s", err.Error())
	}

	engine := api.NewRouter(handler.NewHandler(services))

	err = engine.Run(config.GATEWAY_PORT)
	if err != nil {
		log.Fatalf("error while running server. err: %s", err.Error())
	}
}
