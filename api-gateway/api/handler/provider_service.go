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

// ProviderCreate handles the creation of a new provider.
// @Summary Create provider
// @Description Create a new provider
// @Tags provider
// @Accept json
// @Produce json
// @Param service_id query string false "ServiceId"
// @Success 200 {object} string "Provider created"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /provider/service [post]
func (h *Handler) ProviderServiceCreate(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := claims.(jwt.MapClaims)["user_id"].(string)
	role := claims.(jwt.MapClaims)["role"].(string)

	if role != "provider" {
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}

	var req cp.ProviderServiceRes

	provider_id, err := h.srvs.Provider.GetProviderId(context.Background(), &cp.ById{Id: id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	req.UserId = id
	req.ProviderId = provider_id.Id
	req.ServiceId = c.Query("service_id")

	_, err = h.srvs.ProviderService.Create(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
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
// @Success 200 {object} cp.ProviderServiceGetByIdRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /provider/service/{id} [get]
func (h *Handler) ProviderServiceGetById(c *gin.Context) {
	id := &cp.ById{Id: c.Param("id")}
	res, err := h.srvs.ProviderService.GetById(context.Background(), id)
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
// @Param provider_id query string false "ProviderId"
// @Param page query integer false "Page"
// @Success 200 {object} cp.ProviderServiceGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /provider/service/all [get]
func (h *Handler) ProviderServiceGetAll(c *gin.Context) {
	req := cp.ProviderServiceGetAllReq{
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

	res, err := h.srvs.ProviderService.GetAll(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get provider", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
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
// @Router /provider/service/{id} [delete]
func (h *Handler) ProviderServiceDelete(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	role := claims.(jwt.MapClaims)["role"].(string)

	if role != "provider" {
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}
	
	id := &cp.ById{Id: c.Param("id")}
	_, err := h.srvs.ProviderService.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete provider", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Provider deleted"})
}
