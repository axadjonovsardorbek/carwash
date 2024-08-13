package api

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"

	_ "auth/api/docs"
	"auth/api/handler"
	"auth/api/middleware"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewApi(h *handler.Handler) *gin.Engine {
	router := gin.Default()

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/login", h.Login)
	router.POST("/admin/register", h.AdminRegister)
	router.POST("/user/register", h.UserRegister)
	router.POST("/provider/register", h.ProviderRegister)
	
	
	protected := router.Group("/", middleware.JWTMiddleware())
	protected.GET("/profile", h.Profile)
	protected.GET("/all/users", h.GetAllUsers)
	protected.GET("/refresh-token", h.RefreshToken)
	protected.DELETE("/profile/delete", h.DeleteProfile)
	protected.PUT("/profile/update", h.UpdateProfile)
	protected.PUT("/change-password", h.ChangePassword)
	protected.POST("/reset-password", h.ResetPassword)
	protected.POST("/forgot-password", h.ForgotPassword)

	return router
}
