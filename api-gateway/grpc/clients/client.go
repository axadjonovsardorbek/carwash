package clients

import (
	"gateway/config"
	bp "gateway/genproto/booking"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClients struct {
	Booking         bp.BookingServiceClient
	Notification    bp.NotificationServiceClient
	Payment         bp.PaymentServiceClient
	ProviderService bp.ProviderServiceServiceClient
	Provider        bp.ProviderServiceClient
	Review          bp.ReviewServiceClient
	Service         bp.ServiceServiceClient
}

func NewGrpcClients(cfg *config.Config) (*GrpcClients, error) {
	connB, err := grpc.NewClient(cfg.BOOKING_HOST+cfg.BOOKING_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &GrpcClients{
		Booking:         bp.NewBookingServiceClient(connB),
		Notification:    bp.NewNotificationServiceClient(connB),
		Payment:         bp.NewPaymentServiceClient(connB),
		ProviderService: bp.NewProviderServiceServiceClient(connB),
		Provider:        bp.NewProviderServiceClient(connB),
		Review:          bp.NewReviewServiceClient(connB),
		Service:         bp.NewServiceServiceClient(connB),
	}, nil
}
