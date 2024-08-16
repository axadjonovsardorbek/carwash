package main

import (
	cf "booking/config"
	bp "booking/genproto/booking"
	"booking/kafka"
	"booking/service"
	"booking/storage/mongo"
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

	mongoConn, err := mongo.NewMongoStorage(config)

	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", config.BOOKING_SERVICE_PORT)

	if err != nil {
		log.Fatalf("Failed to listen tcp: %v", err)
	}

	broker := []string{config.KAFKA_HOST + config.KAFKA_PORT}
	kafka, err := kafka.NewKafkaProducer(broker)
	if err != nil {
		log.Fatalln("Failed to connect to Kafka", err)
		return
	}
	defer kafka.Close()

	s := grpc.NewServer()

	bp.RegisterBookingServiceServer(s, service.NewBookingService(db))
	bp.RegisterPaymentServiceServer(s, service.NewPaymentService(db))
	bp.RegisterProviderServiceServer(s, service.NewProviderService(db))
	bp.RegisterReviewServiceServer(s, service.NewReviewService(db))
	bp.RegisterServiceServiceServer(s, service.NewServiceService(db))
	bp.RegisterProviderServiceServiceServer(s, service.NewProviderServiceService(db))
	bp.RegisterNotificationServiceServer(s, service.NewNotificationService(mongoConn))

	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
