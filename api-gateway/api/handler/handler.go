package handler

import (
	"gateway/grpc/clients"
	"gateway/kafka"
)

type Handler struct {
	srvs     *clients.GrpcClients
	Producer kafka.KafkaProducer
}

func NewHandler(srvs *clients.GrpcClients, producer *kafka.KafkaProducer) *Handler {
	return &Handler{
		srvs: srvs,
		Producer: *producer,
	}
}
