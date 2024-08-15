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

// ServiceCreate handles the creation of a new service.
// @Summary Create service
// @Description Create a new service
// @Tags service
// @Accept json
// @Produce json
// @Param service body cp.ServiceCreateReq true "Service data"
// @Success 200 {object} string "Service created"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /service [post]
func (h *Handler) ServiceCreate(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	role := claims.(jwt.MapClaims)["role"].(string)

	if role == "customer" {
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}

	var body cp.ServiceCreateReq
	var req cp.ServiceRes

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req.Name = body.Name
	req.Description = body.Description
	req.Price = body.Price
	req.Duration = body.Duration

	_, err := h.srvs.Service.Create(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Service created"})
}

// ServiceGetById handles the get a service.
// @Summary Get service
// @Description Get a service
// @Tags service
// @Accept json
// @Produce json
// @Param id path string true "Service ID"
// @Success 200 {object} cp.ServiceGetByIdRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /service/{id} [get]
func (h *Handler) ServiceGetById(c *gin.Context) {
	id := &cp.ById{Id: c.Param("id")}
	res, err := h.srvs.Service.GetById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get service", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// ServiceGetAll handles getting all service.
// @Summary Get all service
// @Description Get all service
// @Tags service
// @Accept json
// @Produce json
// @Param page query integer false "Page"
// @Param price query integer false "Price"
// @Success 200 {object} cp.ServiceGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /service/all [get]
func (h *Handler) ServiceGetAll(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	role := claims.(jwt.MapClaims)["role"].(string)

	if role == "customer" {
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}

	req := cp.ServiceGetAllReq{
		Filter: &cp.Filter{},
	}

	priceStr := c.Query("price")
	var price int
	var err error
	if priceStr == "" {
		price = 0
	} else {
		price, err = strconv.Atoi(priceStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price parameter"})
			return
		}
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
	req.Price = int32(price)

	res, err := h.srvs.Service.GetAll(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get service", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// ServiceUpdate handles updating an existing service.
// @Summary Update service
// @Description Update an existing service
// @Tags service
// @Accept json
// @Produce json
// @Param id query string false "Id"
// @Param service body cp.ServiceCreateReq true "Service data"
// @Success 200 {object} string "Service updated"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Service not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /service/{id} [put]
func (h *Handler) ServiceUpdate(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	role := claims.(jwt.MapClaims)["role"].(string)

	if role == "customer" {
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}

	req := cp.ServiceUpdateReq{
		Id:      c.Query("id"),
		Service: &cp.ServiceCreateReq{},
	}

	var body cp.ServiceCreateReq

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req.Service = &body

	_, err := h.srvs.Service.Update(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update service", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Service updated"})
}

// ServiceDelete handles deleting a service by ID.
// @Summary Delete service
// @Description Delete a service by ID
// @Tags service
// @Accept json
// @Produce json
// @Param id path string true "Service ID"
// @Success 200 {object} string "Service deleted"
// @Failure 400 {object} string "Invalid provider ID"
// @Failure 404 {object} string "Service not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /service/{id} [delete]
func (h *Handler) ServiceDelete(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	role := claims.(jwt.MapClaims)["role"].(string)

	if role == "customer" {
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}
	
	id := &cp.ById{Id: c.Param("id")}
	_, err := h.srvs.Service.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete service", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Service deleted"})
}
