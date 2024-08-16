package handler

import (
	"context"
	cp "gateway/genproto/booking"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"google.golang.org/protobuf/encoding/protojson"
)

// BookingCreate handles the creation of a new booking.
// @Summary Create booking
// @Description Create a new booking
// @Tags booking
// @Accept json
// @Produce json
// @Param booking body cp.BookingCreateReq true "Booking data"
// @Success 200 {object} string "Booking created"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /booking [post]
func (h *Handler) BookingCreate(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := claims.(jwt.MapClaims)["user_id"].(string)
	role := claims.(jwt.MapClaims)["role"].(string)

	if role != "customer" {
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}

	var body cp.BookingCreateReq
	var req cp.BookingRes

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	service, err := h.srvs.Service.GetById(context.Background(), &cp.ById{Id: body.ServiceId})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't get service amount"})
		log.Println("error: ", err)
		return
	}

	req.UserId = id
	req.ProviderId = body.ProviderId
	req.ServiceId = body.ServiceId
	req.Status = "pending"
	req.ScheduledTime = body.ScheduledTime
	req.Location = body.Location
	req.TotalPrice = service.Service.Price

	// _, err = h.srvs.Booking.Create(context.Background(), &req)

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	log.Println("error: ", err)
	// 	return
	// }
	data, err := protojson.Marshal(&req)
	if err != nil {
		log.Println("Failed to marshal proto message", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	if err := h.Producer.ProduceMessages("booking-create", data); err != nil {
		log.Println("Failed to produce message to Kafka", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	notification := &cp.NotificationRes{
		UserId:  id,
		Message: "Your book is accepted",
	}
	// _, err = h.srvs.Notification.Create(context.Background(), &cp.NotificationRes{
	// 	UserId:  id,
	// 	Message: "Your book is accepted",
	// })

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	log.Println("error: ", err)
	// 	return
	// }

	data, err = protojson.Marshal(notification)
	if err != nil {
		log.Println("Failed to marshal proto message", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	if err := h.Producer.ProduceMessages("notification-create", data); err != nil {
		log.Println("Failed to produce message to Kafka", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking created"})
}

// BookingGetById handles the get a booking.
// @Summary Get booking
// @Description Get a booking
// @Tags booking
// @Accept json
// @Produce json
// @Param id path string true "Booking ID"
// @Success 200 {object} cp.BookingGetByIdRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /booking/{id} [get]
func (h *Handler) BookingGetById(c *gin.Context) {
	id := &cp.ById{Id: c.Param("id")}
	res, err := h.srvs.Booking.GetById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get booking", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// BookingGetAll handles getting all booking.
// @Summary Get all booking
// @Description Get all booking
// @Tags booking
// @Accept json
// @Produce json
// @Param page query integer false "Page"
// @Success 200 {object} cp.BookingGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /booking/all [get]
func (h *Handler) BookingGetAll(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := claims.(jwt.MapClaims)["user_id"].(string)
	role := claims.(jwt.MapClaims)["role"].(string)

	req := cp.BookingGetAllReq{
		Filter: &cp.Filter{},
	}

	if role == "customer" {
		req.UserId = id
	} else if role == "provider" {
		pr_id, err := h.srvs.Provider.GetProviderId(context.Background(), &cp.ById{Id: id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get provider id", "details": err.Error()})
			return
		}
		req.ProviderId = pr_id.Id
	}

	pageStr := c.Query("page")
	var page int
	var err error
	if pageStr == "" {
		page = 1
	} else {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
			return
		}
	}

	filter := cp.Filter{
		Page: int32(page),
	}

	req.Filter.Page = filter.Page

	res, err := h.srvs.Booking.GetAll(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get bookings", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// BookingUpdate handles updating an existing booking.
// @Summary Update booking
// @Description Update an existing booking
// @Tags booking
// @Accept json
// @Produce json
// @Param id query string false "Id"
// @Success 200 {object} string "Booking updated"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Booking not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /booking/{id} [put]
func (h *Handler) BookingUpdate(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := claims.(jwt.MapClaims)["user_id"].(string)
	role := claims.(jwt.MapClaims)["role"].(string)

	if role != "customer" {
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}

	memory := cp.BookingUpdateReq{
		Id:     c.Query("id"),
		Status: "confirmed",
	}

	// _, err := h.srvs.Booking.Update(context.Background(), &memory)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update booking", "details": err.Error()})
	// 	return
	// }
	data, err := protojson.Marshal(&memory)
	if err != nil {
		log.Println("Failed to marshal proto message", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	if err := h.Producer.ProduceMessages("booking-update", data); err != nil {
		log.Println("Failed to produce message to Kafka", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	notification := &cp.NotificationRes{
		UserId:  id,
		Message: "Your book is confirmed",
	}
	// _, err = h.srvs.Notification.Create(context.Background(), &cp.NotificationRes{
	// 	UserId:  id,
	// 	Message: "Your book is confirmed",
	// })
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	log.Println("error: ", err)
	// 	return
	// }

	data, err = protojson.Marshal(notification)
	if err != nil {
		log.Println("Failed to marshal proto message", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	if err := h.Producer.ProduceMessages("notification-create", data); err != nil {
		log.Println("Failed to produce message to Kafka", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking updated"})
}

// BookingDelete handles deleting a booking by ID.
// @Summary Delete booking
// @Description Delete a booking by ID
// @Tags booking
// @Accept json
// @Produce json
// @Param id path string true "Booking ID"
// @Success 200 {object} string "Booking deleted"
// @Failure 400 {object} string "Invalid booking ID"
// @Failure 404 {object} string "Booking not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /booking/{id} [delete]
func (h *Handler) BookingDelete(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	role := claims.(jwt.MapClaims)["role"].(string)
	user_id := claims.(jwt.MapClaims)["user_id"].(string)

	if role != "customer" {
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}

	id := &cp.ById{Id: c.Param("id")}
	_, err := h.srvs.Booking.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete booking", "details": err.Error()})
		return
	}

	// _, err = h.srvs.Notification.Create(context.Background(), &cp.NotificationRes{
	// 	UserId:  user_id,
	// 	Message: "Your book is cancelled",
	// })
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	log.Println("error: ", err)
	// 	return
	// }
	notification := &cp.NotificationRes{
		UserId:  user_id,
		Message: "Your book is cancelled",
	}

	data, err := protojson.Marshal(notification)
	if err != nil {
		log.Println("Failed to marshal proto message", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	if err := h.Producer.ProduceMessages("notification-create", data); err != nil {
		log.Println("Failed to produce message to Kafka", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking deleted"})
}
