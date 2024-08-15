package handler

import (
	"context"
	cp "gateway/genproto/booking"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// PaymentCreate handles the creation of a new payment.
// @Summary Create payment
// @Description Create a new payment
// @Tags payment
// @Accept json
// @Produce json
// @Param payment body cp.PaymentCreateReq true "Payment data"
// @Success 200 {object} string "Payment created"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /payment [post]
func (h *Handler) PaymentCreate(c *gin.Context) {
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

	var body cp.PaymentCreateReq
	var req cp.PaymentRes

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	booking_id, err := h.srvs.Booking.GetById(context.Background(), &cp.ById{Id: body.BookingId})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	amount, err := h.srvs.Payment.GetBookingAmount(context.Background(), &cp.ById{Id: booking_id.Booking.ServiceId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}


	req.UserId = id
	req.BookingId = body.BookingId
	req.Amount = amount.Amount
	req.PaymentMethod = body.PaymentMethod
	req.Status = "pending"

	_, err = h.srvs.Payment.Create(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	_, err = h.srvs.Notification.Create(context.Background(), &cp.NotificationRes{
		UserId: id,
		Message: "Your payment is accepted",
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment created"})
}

// PaymentGetById handles the get a payment.
// @Summary Get payment
// @Description Get a payment
// @Tags payment
// @Accept json
// @Produce json
// @Param id path string true "Payment ID"
// @Success 200 {object} cp.PaymentGetByIdRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /payment/{id} [get]
func (h *Handler) PaymentGetById(c *gin.Context) {
	id := &cp.ById{Id: c.Param("id")}
	res, err := h.srvs.Payment.GetById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get payment", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// PaymentGetAll handles getting all payment.
// @Summary Get all payment
// @Description Get all payment
// @Tags payment
// @Accept json
// @Produce json
// @Param status query string false "Status"
// @Param page query integer false "Page"
// @Success 200 {object} cp.PaymentGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /payment/all [get]
func (h *Handler) PaymentGetAll(c *gin.Context) {
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

	req := cp.PaymentGetAllReq{
		UserId: id,
		Status: c.Query("status"),
		Filter: &cp.Filter{},
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

	res, err := h.srvs.Payment.GetAll(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get payment", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// PaymentUpdate handles updating an existing payment.
// @Summary Update payment
// @Description Update an existing payment
// @Tags payment
// @Accept json
// @Produce json
// @Param id query string false "Id"
// @Success 200 {object} string "Payment updated"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Payment not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /payment/{id} [put]
func (h *Handler) PaymentUpdate(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user_id := claims.(jwt.MapClaims)["user_id"].(string)
	role := claims.(jwt.MapClaims)["role"].(string)

	if role != "customer" {
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}

	payment_id := c.Query("id")

	memory := cp.PaymentUpdateReq{
		Id:     payment_id,
		Status: "completed",
	}

	_, err := h.srvs.Payment.Update(context.Background(), &memory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update payment", "details": err.Error()})
		return
	}

	_, err = h.srvs.Notification.Create(context.Background(), &cp.NotificationRes{
		UserId: user_id,
		Message: "Your payment is completed",
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	booking_id, err := h.srvs.Payment.GetBookingId(context.Background(), &cp.ById{
		Id: payment_id,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	_, err = h.srvs.Booking.Update(context.Background(), &cp.BookingUpdateReq{
		Id: booking_id.Id,
		Status: "completed",
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	_, err = h.srvs.Notification.Create(context.Background(), &cp.NotificationRes{
		UserId: user_id,
		Message: "Your book is completed",
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment updated"})
}

// PaymentDelete handles deleting a payment by ID.
// @Summary Delete payment
// @Description Delete a payment by ID
// @Tags payment
// @Accept json
// @Produce json
// @Param id path string true "Payment ID"
// @Success 200 {object} string "Payment cancelled"
// @Failure 400 {object} string "Invalid payment ID"
// @Failure 404 {object} string "Payment not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /payment/{id} [delete]
func (h *Handler) PaymentDelete(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user_id := claims.(jwt.MapClaims)["user_id"].(string)
	role := claims.(jwt.MapClaims)["role"].(string)

	if role != "customer" {
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}

	id := &cp.ById{Id: c.Param("id")}
	_, err := h.srvs.Payment.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't cancel payment", "details": err.Error()})
		return
	}

	_, err = h.srvs.Notification.Create(context.Background(), &cp.NotificationRes{
		UserId: user_id,
		Message: "Your payment is cancelled",
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment cancelled"})
}
