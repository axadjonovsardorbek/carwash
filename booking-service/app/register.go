package app

import (
	"errors"

	"booking/config"
	"booking/kafka"
)

func Register(h *KafkaHandler, cfg *config.Config) error {

	brokers := []string{cfg.KAFKA_HOST + cfg.KAFKA_PORT}
	kcm := kafka.NewKafkaConsumerManager()

	if err := kcm.RegisterConsumer(brokers, "booking-create", "booking-create-id", h.BookingCreateHandler()); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'booking-create' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "booking-update", "booking-update-id", h.BookingUpdateHandler()); err!= nil {
		if err == kafka.ErrConsumerAlreadyExists {
            return errors.New("consumer for topic 'booking-update' already exists")
        } else {
            return errors.New("error registering consumer:" + err.Error())
        }
	}
	if err := kcm.RegisterConsumer(brokers, "notification-create", "notification-create-id", h.NotificationCreateHandler()); err!= nil {
		if err == kafka.ErrConsumerAlreadyExists {
            return errors.New("consumer for topic 'notification-create' already exists")
        } else {
            return errors.New("error registering consumer:" + err.Error())
        }
	}
	if err := kcm.RegisterConsumer(brokers, "notification-update", "notification-update-id", h.NotificationUpdateHandler()); err!= nil {
		if err == kafka.ErrConsumerAlreadyExists {
            return errors.New("consumer for topic 'notification-update' already exists")
        } else {
            return errors.New("error registering consumer:" + err.Error())
        }
	}
	if err := kcm.RegisterConsumer(brokers, "payment-create", "payment-create-id", h.PaymentCreateHandler()); err!= nil {
		if err == kafka.ErrConsumerAlreadyExists {
            return errors.New("consumer for topic 'payment-create' already exists")
        } else {
            return errors.New("error registering consumer:" + err.Error())
        }
	}
	if err := kcm.RegisterConsumer(brokers, "payment-update", "payment-update-id", h.PaymentUpdateHandler()); err!= nil {
		if err == kafka.ErrConsumerAlreadyExists {
            return errors.New("consumer for topic 'payment-update' already exists")
        } else {
            return errors.New("error registering consumer:" + err.Error())
        }
	}
	if err := kcm.RegisterConsumer(brokers, "provider-create", "provider-create-id", h.PaymentCreateHandler()); err!= nil {
		if err == kafka.ErrConsumerAlreadyExists {
            return errors.New("consumer for topic 'provider-create' already exists")
        } else {
            return errors.New("error registering consumer:" + err.Error())
        }
	}
	if err := kcm.RegisterConsumer(brokers, "provider-update", "provider-update-id", h.PaymentUpdateHandler()); err!= nil {
		if err == kafka.ErrConsumerAlreadyExists {
            return errors.New("consumer for topic 'provider-update' already exists")
        } else {
            return errors.New("error registering consumer:" + err.Error())
        }
	}
	if err := kcm.RegisterConsumer(brokers, "review-create", "review-create-id", h.ReviewCreateHandler()); err!= nil {
		if err == kafka.ErrConsumerAlreadyExists {
            return errors.New("consumer for topic 'review-create' already exists")
        } else {
            return errors.New("error registering consumer:" + err.Error())
        }
	}
	if err := kcm.RegisterConsumer(brokers, "review-update", "review-update-id", h.ReviewUpdateHandler()); err!= nil {
		if err == kafka.ErrConsumerAlreadyExists {
            return errors.New("consumer for topic 'review-update' already exists")
        } else {
            return errors.New("error registering consumer:" + err.Error())
        }
	}
	return nil

}
