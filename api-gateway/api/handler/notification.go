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

// NotificationGetById handles the get a notification.
// @Summary Get notification
// @Description Get a notification
// @Tags notification
// @Accept json
// @Produce json
// @Param id path string true "Notification ID"
// @Success 200 {object} cp.NotificationGetByIdRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /notification/{id}/read [get]
func (h *Handler) NotificationGetById(c *gin.Context) {
	id := &cp.ById{Id: c.Param("id")}
	res, err := h.srvs.Notification.GetById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get notification", "details": err.Error()})
		return
	}

	notification := &cp.NotificationUpdateReq{
		IsRead: "true",
		Id:     id.Id,
	}

	data, err := protojson.Marshal(notification)
	if err != nil {
		log.Println("Failed to marshal proto message", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	if err := h.Producer.ProduceMessages("notification-update", data); err != nil {
		log.Println("Failed to produce message to Kafka", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

// NotificationGetAll handles getting all notification.
// @Summary Get all notification
// @Description Get all notification
// @Tags notification
// @Accept json
// @Produce json
// @Param page query integer false "Page"
// @Param is_read query string false "IsRead"
// @Success 200 {object} cp.NotificationGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /notification/all [get]
func (h *Handler) NotificationGetAll(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := claims.(jwt.MapClaims)["user_id"].(string)

	req := cp.NotificationGetAllReq{
		UserId: id,
		IsRead: c.Query("is_read"),
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

	res, err := h.srvs.Notification.GetAll(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get notification", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
