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

// ProviderCreate handles the creation of a new provider.
// @Summary Create provider
// @Description Create a new provider
// @Tags provider
// @Accept json
// @Produce json
// @Param provider body cp.ProviderCreateReq true "Provider data"
// @Success 200 {object} string "Provider created"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /provider [post]
func (h *Handler) ProviderCreate(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := claims.(jwt.MapClaims)["user_id"].(string)
	role := claims.(jwt.MapClaims)["role"].(string)
	email := claims.(jwt.MapClaims)["email"].(string)

	if role != "provider" {
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}

	var body cp.ProviderCreateReq
	var req cp.ProviderRes

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req.UserId = id
	req.CompanyName = body.CompanyName
	req.Description = body.Description
	req.Availability = body.Availability
	req.AverageRating = 0
	req.Location = body.Location

	data, err := protojson.Marshal(&req)
	if err != nil {
		log.Println("Failed to marshal proto message", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	if err := h.Producer.ProduceMessages("provider-create", data); err != nil {
		log.Println("Failed to produce message to Kafka", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	notification := &cp.NotificationRes{
		UserId:  id,
		Message: "Your provider is added",
		Email:   email,
	}
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

	c.JSON(http.StatusOK, gin.H{"message": "Provider created"})
}

// ProviderGetById handles the get a provider.
// @Summary Get provider
// @Description Get a provider
// @Tags provider
// @Accept json
// @Produce json
// @Param id path string true "Provider ID"
// @Success 200 {object} cp.ProviderGetByIdRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /provider/{id} [get]
func (h *Handler) ProviderGetById(c *gin.Context) {
	id := &cp.ById{Id: c.Param("id")}
	res, err := h.srvs.Provider.GetById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get provider", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// ProviderGetAll handles getting all provider.
// @Summary Get all provider
// @Description Get all provider
// @Tags provider
// @Accept json
// @Produce json
// @Param page query integer false "Page"
// @Param rating query number false "Rating"
// @Success 200 {object} cp.ProviderGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /provider/all [get]
func (h *Handler) ProviderGetAll(c *gin.Context) {
	ratingStr := c.Query("rating")
	var rating float64
	var err error
	if ratingStr == "" {
		rating = 1
	} else {
		rating, err = strconv.ParseFloat(ratingStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rating"})
			return
		}
	}

	req := cp.ProviderGetAllReq{
		AverageRating: float32(rating),
		Filter:        &cp.Filter{},
	}

	pageStr := c.Query("page")
	var page int
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

	res, err := h.srvs.Provider.GetAll(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get provider", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// ProviderUpdate handles updating an existing provider.
// @Summary Update provider
// @Description Update an existing provider
// @Tags provider
// @Accept json
// @Produce json
// @Param id query string false "Id"
// @Param provider body cp.ProviderCreateReq true "Provider data"
// @Success 200 {object} string "Provider updated"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Provider not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /provider/{id} [put]
func (h *Handler) ProviderUpdate(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	role := claims.(jwt.MapClaims)["role"].(string)
	user_id := claims.(jwt.MapClaims)["user_id"].(string)
	email := claims.(jwt.MapClaims)["email"].(string)

	if role != "provider" {
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}

	req := cp.ProviderUpdateReq{
		Id:       c.Query("id"),
		Provider: &cp.ProviderCreateReq{},
	}

	var body cp.ProviderCreateReq

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	pr_rating, err := h.srvs.Review.GetProviderRating(context.Background(), &cp.ById{Id: req.Id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get rating", "details": err.Error()})
		return
	}

	avg_rating := float32(pr_rating.Rating) / float32(pr_rating.Count)

	req.Provider.UserId = user_id
	req.Provider.CompanyName = body.CompanyName
	req.Provider.Description = body.Description
	req.Provider.Availability = body.Availability
	req.Provider.AverageRating = avg_rating
	req.Provider.Location = body.Location

	data, err := protojson.Marshal(&req)
	if err != nil {
		log.Println("Failed to marshal proto message", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	if err := h.Producer.ProduceMessages("provider-update", data); err != nil {
		log.Println("Failed to produce message to Kafka", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	notification := &cp.NotificationRes{
		UserId:  user_id,
		Message: "Your provider is updated",
		Email:   email,
	}
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

	c.JSON(http.StatusOK, gin.H{"message": "Provider updated"})
}

// ProviderDelete handles deleting a provider by ID.
// @Summary Delete provider
// @Description Delete a provider by ID
// @Tags provider
// @Accept json
// @Produce json
// @Param id path string true "Provider ID"
// @Success 200 {object} string "Provider deleted"
// @Failure 400 {object} string "Invalid provider ID"
// @Failure 404 {object} string "Provider not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /provider/{id} [delete]
func (h *Handler) ProviderDelete(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	role := claims.(jwt.MapClaims)["role"].(string)
	user_id := claims.(jwt.MapClaims)["user_id"].(string)
	email := claims.(jwt.MapClaims)["email"].(string)

	if role != "provider" {
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}

	id := &cp.ById{Id: c.Param("id")}
	_, err := h.srvs.Provider.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete provider", "details": err.Error()})
		return
	}

	notification := &cp.NotificationRes{
		UserId:  user_id,
		Message: "Your provider is deleted",
		Email:   email,
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

	c.JSON(http.StatusOK, gin.H{"message": "Provider deleted"})
}
