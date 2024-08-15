package mongo

import (
	"booking/config"
	"booking/storage"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	Db        *mongo.Database
	// ProviderS storage.ProviderI
	// ServiceS  storage.ServiceI
	// BookingS  storage.BookingI
	// ReviewS   storage.ReviewI
	// PaymentS  storage.PaymentI
	NotificationS storage.NotificationI
	// ProviderServiceS storage.ProviderServiceI
}

func NewMongoStorage(config config.Config) (*Storage, error) {
	uri := fmt.Sprintf("mongodb://%s:%d", config.MONGO_DB_HOST, config.MONGO_DB_PORT)

	clientOptions := options.Client().ApplyURI(uri).SetAuth(options.Credential{Username: config.MONGO_DB_USER, Password: config.MONGO_DB_PASS})

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(config.MONGO_DB_NAME)

	// provider := NewProviderRepo(db)
	// service := NewServiceRepo(db)
	// booking := NewBookingRepo(db)
	// review := NewReviewRepo(db)
	// payment := NewPaymentRepo(db)
	notification := NewNotificationRepo(db)
	// provider_service := NewProviderServiceRepo(db)

	fmt.Println("Connected to MongoDB!")
	return &Storage{
		Db:        db,
		// ProviderS: provider,
		// ServiceS:  service,
		// BookingS:  booking,
		// ReviewS:   review,
		// PaymentS:  payment,
		NotificationS:  notification,
		// ProviderServiceS: provider_service,
	}, nil
}

// func (s *Storage) Provider() storage.ProviderI {
// 	if s.ProviderS == nil {
// 		s.ProviderS = NewProviderRepo(s.Db)
// 	}
// 	return s.ProviderS
// }

// func (s *Storage) ProviderService() storage.ProviderServiceI {
// 	if s.ProviderServiceS == nil {
// 		s.ProviderServiceS = NewProviderServiceRepo(s.Db)
// 	}
// 	return s.ProviderServiceS
// }

// func (s *Storage) Service() storage.ServiceI {
// 	if s.ServiceS == nil {
// 		s.ServiceS = NewServiceRepo(s.Db)
// 	}
// 	return s.ServiceS
// }

// func (s *Storage) Booking() storage.BookingI {
// 	if s.BookingS == nil {
// 		s.BookingS = NewBookingRepo(s.Db)
// 	}
// 	return s.BookingS
// }

// func (s *Storage) Review() storage.ReviewI {
// 	if s.ReviewS == nil {
// 		s.ReviewS = NewReviewRepo(s.Db)
// 	}
// 	return s.ReviewS
// }

func (s *Storage) Notification() storage.NotificationI {
	if s.NotificationS == nil {
		s.NotificationS = NewNotificationRepo(s.Db)
	}
	return s.NotificationS
}

// func (s *Storage) Payment() storage.PaymentI {
// 	if s.PaymentS == nil {
// 		s.PaymentS = NewPaymentRepo(s.Db)
// 	}
// 	return s.PaymentS
// }
