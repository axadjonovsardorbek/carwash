package main

import (
	cf "booking/config"
	bp "booking/genproto/booking"
	"booking/service"
	"booking/storage/postgres"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	config := cf.Load()

	db, err := postgres.NewPostgresStorage(config)

	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", config.BOOKING_SERVICE_PORT)

	if err != nil {
		log.Fatalf("Failed to listen tcp: %v", err)
	}

	s := grpc.NewServer()

	bp.RegisterBookingServiceServer(s, service.NewBookingService(db))
	bp.RegisterPaymentServiceServer(s, service.NewPaymentService(db))
	bp.RegisterProviderServiceServer(s, service.NewProviderService(db))
	bp.RegisterReviewServiceServer(s, service.NewReviewService(db))
	bp.RegisterServiceServiceServer(s, service.NewServiceService(db))
	bp.RegisterProviderServiceServiceServer(s, service.NewProviderServiceService(db))

	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
