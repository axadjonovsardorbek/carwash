package handler

import (
	"auth/kafka"
	"auth/service"
)

type Handler struct {
	User     *service.UsersService
	Producer kafka.KafkaProducer
}

func NewHandler(us *service.UsersService, kafka kafka.KafkaProducer) *Handler {
	return &Handler{User: us, Producer: kafka}
}
