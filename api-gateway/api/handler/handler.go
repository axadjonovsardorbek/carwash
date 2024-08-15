package handler

import (
	"gateway/grpc/clients"
)

type Handler struct {
	srvs *clients.GrpcClients
}

func NewHandler(srvs *clients.GrpcClients) *Handler {
	return &Handler{srvs: srvs}
}