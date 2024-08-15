package api

import (
	"gateway/api/handler"
	"gateway/api/middleware"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	_ "gateway/api/docs"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewRouter(h *handler.Handler) *gin.Engine {

	router := gin.Default()

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	protected := router.Group("/", middleware.JWTMiddleware())
	
	booking := protected.Group("/booking")
	booking.POST("/", h.BookingCreate)
	booking.GET("/:id", h.BookingGetById)
	booking.GET("/all", h.BookingGetAll)
	booking.PUT("/:id", h.BookingUpdate)
	booking.DELETE("/:id", h.BookingDelete)

	provider := protected.Group("/provider")
	provider.POST("/", h.ProviderCreate)
	provider.GET("/:id", h.ProviderGetById)
	provider.GET("/all", h.ProviderGetAll)
	provider.PUT("/:id", h.ProviderUpdate)
	provider.DELETE("/:id", h.ProviderDelete)

	p_service := provider.Group("/service")
	p_service.POST("/", h.ProviderServiceCreate)
	p_service.GET("/:id", h.ProviderServiceGetById)
	p_service.GET("/all", h.ProviderServiceGetAll)
	p_service.DELETE("/:id", h.ProviderServiceDelete)

	review := provider.Group("/review")
	review.POST("/", h.ReviewCreate)
	review.GET("/:id", h.ReviewGetById)
	review.GET("/all", h.ReviewGetAll)
	review.PUT("/:id", h.ReviewUpdate)
	// review.DELETE("/:id", h.ReviewDelete)

	payment := protected.Group("/payment")
	payment.POST("/", h.PaymentCreate)
	payment.GET("/:id", h.PaymentGetById)
	payment.GET("/all", h.PaymentGetAll)
	payment.PUT("/:id", h.PaymentUpdate)
	payment.DELETE("/:id", h.PaymentDelete)

	service := protected.Group("/service")
	service.POST("/", h.ServiceCreate)
	service.GET("/:id", h.ServiceGetById)
	service.GET("/all", h.ServiceGetAll)
	service.PUT("/:id", h.ServiceUpdate)
	service.DELETE("/:id", h.ServiceDelete)

	notification := protected.Group("/notification")
	notification.GET("/:id/read", h.NotificationGetById)
	notification.GET("/all", h.NotificationGetAll)

	return router
}
