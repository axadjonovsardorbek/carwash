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

// ReviewCreate handles the creation of a new review.
// @Summary Create review
// @Description Create a new review
// @Tags review
// @Accept json
// @Produce json
// @Param review body cp.ReviewCreateReq true "Review data"
// @Success 200 {object} string "Review created"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /provider/review [post]
func (h *Handler) ReviewCreate(c *gin.Context) {
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

	var body cp.ReviewCreateReq
	var req cp.ReviewRes

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	provider_id, err := h.srvs.Booking.GetById(context.Background(), &cp.ById{Id: body.BookingId})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	if body.Rating < 1 || body.Rating > 5 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "you aren't choose rating point"})
		return
	}

	req.BookingId = body.BookingId
	req.UserId = id
	req.ProviderId = provider_id.Booking.ProviderId
	req.Rating = body.Rating
	req.Comment = body.Comment

	data, err := protojson.Marshal(&req)
	if err != nil {
		log.Println("Failed to marshal proto message", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	if err := h.Producer.ProduceMessages("review-create", data); err != nil {
		log.Println("Failed to produce message to Kafka", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	pr_rating, err := h.srvs.Review.GetProviderRating(context.Background(), &cp.ById{Id: req.ProviderId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get rating", "details": err.Error()})
		return
	}

	avg_rating := float32(pr_rating.Rating) / float32(pr_rating.Count)

	_, err = h.srvs.Provider.Update(context.Background(), &cp.ProviderUpdateReq{
		Id: req.ProviderId,
		Provider: &cp.ProviderCreateReq{
			AverageRating: avg_rating,
		},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review created"})
}

// ReviewGetById handles the get a review.
// @Summary Get review
// @Description Get a review
// @Tags review
// @Accept json
// @Produce json
// @Param id path string true "Review ID"
// @Success 200 {object} cp.ReviewGetByIdRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /provider/review/{id} [get]
func (h *Handler) ReviewGetById(c *gin.Context) {
	id := &cp.ById{Id: c.Param("id")}
	res, err := h.srvs.Review.GetById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get review", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// ReviewGetAll handles getting all review.
// @Summary Get all review
// @Description Get all review
// @Tags review
// @Accept json
// @Produce json
// @Param page query integer false "Page"
// @Param provider_id query string false "ProviderId"
// @Param user_id query string false "UserId"
// @Success 200 {object} cp.ReviewGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /provider/review/all [get]
func (h *Handler) ReviewGetAll(c *gin.Context) {
	req := cp.ReviewGetAllReq{
		ProviderId: c.Query("provider_id"),
		Filter:     &cp.Filter{},
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

	res, err := h.srvs.Review.GetAll(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get review", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// ReviewUpdate handles updating an existing review.
// @Summary Update review
// @Description Update an existing review
// @Tags review
// @Accept json
// @Produce json
// @Param id query string false "Id"
// @Param comment query string false "Comment"
// @Param rating query integer false "Rating"
// @Success 200 {object} string "Review updated"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Review not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /provider/review/{id} [put]
func (h *Handler) ReviewUpdate(c *gin.Context) {
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

	ratingStr := c.Query("rating")
	var rating float64
	var err error
	if ratingStr == "" {
		rating = 0
	} else {
		rating, err = strconv.ParseFloat(ratingStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rating"})
			return
		}
	}

	req := cp.ReviewUpdateReq{
		Id:      c.Query("id"),
		Comment: c.Query("comment"),
		Rating:  int32(rating),
		UserId:  id,
	}

	if req.Rating != 0 && req.Rating < 1 || req.Rating > 5 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "you aren't choose rating point"})
		return
	}

	_, err = h.srvs.Review.Update(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update review", "details": err.Error()})
		return
	}

	provider_id, err := h.srvs.Review.GetById(context.Background(), &cp.ById{Id: req.Id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get review", "details": err.Error()})
		return
	}

	pr_rating, err := h.srvs.Review.GetProviderRating(context.Background(), &cp.ById{Id: provider_id.Review.ProviderId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get rating", "details": err.Error()})
		return
	}

	avg_rating := float32(pr_rating.Rating) / float32(pr_rating.Count)

	_, err = h.srvs.Provider.Update(context.Background(), &cp.ProviderUpdateReq{
		Id: provider_id.Review.ProviderId,
		Provider: &cp.ProviderCreateReq{
			AverageRating: avg_rating,
		}})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update provider", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review updated"})
}
