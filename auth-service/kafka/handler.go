package kafka

import (
	"context"
	"encoding/json"
	"log"

	ap "auth/genproto/auth"
	"auth/service"
)

func UserCreateHandler(userService *service.UsersService) func(message []byte) {
	return func(message []byte) {
		var user ap.UserCreateReq
		if err := json.Unmarshal(message, &user); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := userService.Register(context.Background(), &user)
		if err != nil {
			log.Printf("Cannot create user via Kafka: %v", err)
			return
		}
		log.Printf("Created user: %+v", res)
	}
}

func ForgotPasswordHandler(userService *service.UsersService) func(message []byte) {
	return func(message []byte) {
        var email ap.UsersForgotPassword
        if err := json.Unmarshal(message, &email); err!= nil {
            log.Printf("Cannot unmarshal JSON: %v", err)
            return
        }

        res, err := userService.ForgotPassword(context.Background(), &email)
        if err!= nil {
            log.Printf("Cannot send forgot password via Kafka: %v", err)
            return
        }
        log.Printf("Sent forgot password email: %+v", res)
    }
}