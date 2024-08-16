package app

import (
	"booking/config"
	"booking/kafka"
	"booking/service"
	"booking/storage/mongo"
	"booking/storage/postgres"
	"log"
	"net"

	bp "booking/genproto/booking"

	"google.golang.org/grpc"
)

func Run(cfg config.Config) {

	db, err := postgres.NewPostgresStorage(cfg)

	if err != nil {
		panic(err)
	}

	mongoConn, err := mongo.NewMongoStorage(cfg)

	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", cfg.BOOKING_SERVICE_PORT)

	if err != nil {
		log.Fatalf("Failed to listen tcp: %v", err)
	}

	kafka, err := kafka.NewKafkaProducer([]string{cfg.KAFKA_HOST + cfg.KAFKA_PORT})
	if err != nil {
		log.Fatal(err)
	}

	// create services

	booking_service := service.NewBookingService(db)
	notification_service := service.NewNotificationService(mongoConn)
	payment_service := service.NewPaymentService(db)
	provider_service := service.NewProviderService(db)
	review_service := service.NewReviewService(db)
	// service_service := service.NewServiceService(db)

	//register kafka handlers
	kafka_handler := &KafkaHandler{
		booking:      booking_service,
		notification: notification_service,
		payment:      payment_service,
		provider:     provider_service,
		review:       review_service,
	}

	// register kafka handlers
	if err := Register(kafka_handler, &cfg); err != nil {
		log.Fatal("Error registering kafka handlers: ", err)
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
