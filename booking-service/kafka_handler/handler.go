package kafka_handler

import (
	"context"
	"log"

	pb "booking/genproto/booking"
	"booking/service"

	"google.golang.org/protobuf/encoding/protojson"
)

type KafkaHandler struct {
	booking      *service.BookingService
	notification *service.NotificationService
	payment      *service.PaymentService
	provider     *service.ProviderService
	// provider_service        *service.ProviderServiceService
	review *service.ReviewService
	// service *service.ServiceService
}

func (h *KafkaHandler) BookingCreateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.BookingRes
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.booking.Create(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot create booking via Kafka: %v", err)
			return
		}
		log.Printf("Created booking: %+v", res)
	}
}
func (h *KafkaHandler) BookingUpdateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.BookingUpdateReq
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.booking.Update(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot update booking via Kafka: %v", err)
			return
		}
		log.Printf("Updated booking: %+v", res)
	}
}
func (h *KafkaHandler) NotificationCreateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.NotificationRes
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.notification.Create(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot create notification via Kafka: %v", err)
			return
		}
		log.Printf("Created notification: %+v", res)
	}
}
func (h *KafkaHandler) NotificationUpdateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.NotificationUpdateReq
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.notification.Update(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot update notification via Kafka: %v", err)
			return
		}
		log.Printf("Updated notification: %+v", res)
	}
}
func (h *KafkaHandler) PaymentCreateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.PaymentRes
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.payment.Create(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot create payment via Kafka: %v", err)
			return
		}
		log.Printf("Created payment: %+v", res)
	}
}
func (h *KafkaHandler) PaymentUpdateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.PaymentUpdateReq
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.payment.Update(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot update payment via Kafka: %v", err)
			return
		}
		log.Printf("Updated payment: %+v", res)
	}
}
func (h *KafkaHandler) ProviderCreateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.ProviderRes
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.provider.Create(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot create provider via Kafka: %v", err)
			return
		}
		log.Printf("Created provider: %+v", res)
	}
}
func (h *KafkaHandler) ProviderUpdateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.ProviderUpdateReq
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.provider.Update(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot update provider via Kafka: %v", err)
			return
		}
		log.Printf("Updated provider: %+v", res)
	}
}
func (h *KafkaHandler) ReviewCreateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.ReviewRes
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.review.Create(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot create review via Kafka: %v", err)
			return
		}
		log.Printf("Created review: %+v", res)
	}
}
func (h *KafkaHandler) ReviewUpdateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.ReviewUpdateReq
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.review.Update(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot update review via Kafka: %v", err)
			return
		}
		log.Printf("Updated review: %+v", res)
	}
}
