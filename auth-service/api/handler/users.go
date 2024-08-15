package handler

import (
	ap "auth/genproto/auth"
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/golang-jwt/jwt"
	_ "github.com/swaggo/swag"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

const emailRegex = `^[a-zA-Z0-9._]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

func isValidEmail(email string) bool {
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

// Register godoc
// @Summary Register a new admin
// @Description Register a new admin
// @Tags admin
// @Accept json
// @Produce json
// @Param user body ap.UserCreateReq true "User registration request"
// @Success 201 {object} string "Admin registered"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Router /admin/register [post]
func (h *Handler) AdminRegister(c *gin.Context) {
	var req ap.UserCreateReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !isValidEmail(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect email"})
		return
	}

	req.Role = "admin"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	req.Password = string(hashedPassword)

	_, err = h.User.Register(context.Background(), &req)

	// input, err := json.Marshal(&req)
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error()})
	// 	return
	// }
	// err = h.Producer.ProduceMessages("create", input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Admin registered"})
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body ap.UserCreateReq true "User registration request"
// @Success 201 {object} string "User registered"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Router /user/register [post]
func (h *Handler) UserRegister(c *gin.Context) {
	var req ap.UserCreateReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !isValidEmail(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect email"})
		return
	}

	req.Role = "customer"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	req.Password = string(hashedPassword)

	_, err = h.User.Register(context.Background(), &req)

	// input, err := json.Marshal(&req)
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error()})
	// 	return
	// }
	// err = h.Producer.ProduceMessages("create", input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user_id, err := h.User.CheckEmail(context.Background(), &ap.CheckEmailReq{Email: req.Email})

	if err != nil || user_id.Id == "" || user_id.Id == "string" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		log.Println("error: ", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered"})
}

// Register godoc
// @Summary Register a new courier
// @Description Register a new courier
// @Tags provider
// @Accept json
// @Produce json
// @Param user body ap.UserCreateReq true "User registration request"
// @Success 201 {object} string "Courier registered"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Router /provider/register [post]
func (h *Handler) ProviderRegister(c *gin.Context) {
	var req ap.UserCreateReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !isValidEmail(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect email"})
		return
	}

	req.Role = "provider"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	req.Password = string(hashedPassword)

	_, err = h.User.Register(context.Background(), &req)

	// input, err := json.Marshal(&req)
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error()})
	// 	return
	// }
	// err = h.Producer.ProduceMessages("create", input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Courier registered"})
}

// Login godoc
// @Summary Login a user
// @Description Authenticate user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body ap.UserLoginReq true "User login credentials"
// @Success 200 {object} ap.TokenRes "JWT tokens"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Invalid email or password"
// @Router /login [post]
func (h *Handler) Login(c *gin.Context) {
	var req ap.UserLoginReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.User.Login(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetProfile godoc
// @Summary Get user profile
// @Description Get the profile of the authenticated user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} ap.UserRes
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "User not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /profile [get]
func (h *Handler) Profile(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := claims.(jwt.MapClaims)["user_id"].(string)

	user, err := h.User.Profile(context.Background(), &ap.ById{Id: id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUsers godoc
// @Summary Get users
// @Description Get the profile of the authenticated users
// @Tags user
// @Accept json
// @Produce json
// @Param role query string false "Role"
// @Param page query integer false "Page"
// @Success 200 {object} ap.GetAllUsersRes
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "User not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /all/users [get]
func (h *Handler) GetAllUsers(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	role := claims.(jwt.MapClaims)["role"].(string)

	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}

	req := ap.GetAllUsersReq{
		Role: c.Query("role"),
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

	req.Page = int32(page)

	res, err := h.User.GetAllUsers(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get users", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateProfile godoc
// @Summary Update user profile
// @Description Update the profil of the authenticated user
// @Tags user
// @Accept json
// @Produce json
// @Param first_name query string false "FirstName"
// @Param last_name query string false "LastName"
// @Param email query string false "Email"
// @Success 200 {object} string "User profile updated"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "User settings not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /profile/update [put]
func (h *Handler) UpdateProfile(c *gin.Context) {
	var req ap.UserUpdateReq

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := claims.(jwt.MapClaims)["user_id"].(string)
	first_name := c.Query("first_name")
	last_name := c.Query("last_name")
	email := c.Query("email")

	if !isValidEmail(email) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect email"})
		return
	}

	req.Id = id
	req.FirstName = first_name
	req.LastName = last_name
	req.Email = email

	_, err := h.User.UpdateProfile(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User profile updated"})
}

// DeleteProfile godoc
// @Summary Delete user profile
// @Description Delete the profil of the authenticated user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} string "User profile deleted"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "User not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /profile/delete [delete]
func (h *Handler) DeleteProfile(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := claims.(jwt.MapClaims)["user_id"].(string)

	_, err := h.User.DeleteProfile(context.Background(), &ap.ById{Id: id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User profile deleted"})
}

// ChangePassword godoc
// @Summary ChangePassword
// @Description ChangePassword
// @Tags user
// @Accept json
// @Produce json
// @Param current_password query string false "CurrentPassword"
// @Param new_password query string false "NewPassword"
// @Success 200 {object} string "Changed password"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Password incorrect"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /change-password [put]
func (h *Handler) ChangePassword(c *gin.Context) {
	var req ap.UsersChangePassword

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := claims.(jwt.MapClaims)["user_id"].(string)
	cur_pass := c.Query("current_password")
	new_pass := c.Query("new_password")

	if cur_pass == "" || cur_pass == "string" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Password incorrect"})
		return
	}
	if new_pass == "" || new_pass == "string" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Password incorrect"})
		return
	}

	req.Id = id
	req.CurrentPassword = cur_pass
	req.NewPassword = new_pass

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	req.NewPassword = string(hashedPassword)

	_, err = h.User.ChangePassword(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Changed password"})
}

// ForgotPassword godoc
// @Summary Send a reset password code to the user's email
// @Description Send a reset password code to the user's email
// @Tags user
// @Accept  json
// @Produce  json
// @Param  email  body  ap.UsersForgotPassword  true  "Email data"
// @Success 200 {object} string "Reset password code sent successfully"
// @Failure 400 {object} string "Invalid input"
// @Failure 500 {object} string "Internal server error"
// @Security BearerAuth
// @Router /forgot-password [post]
func (h *Handler) ForgotPassword(c *gin.Context) {
	req := ap.UsersForgotPassword{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := h.User.ForgotPassword(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(req.Email)

	// input,err := json.Marshal(req)
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error()})
	// }

	// err = h.Producer.ProduceMessages("forgot_password",input)
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(200, gin.H{"message": "Reset password code sent successfully"})
}

// ResetPassword godoc
// @Summary Reset user password
// @Description Reset user password with the provided reset code and new password
// @Tags user
// @Accept  json
// @Produce  json
// @Param reset_token query string false "ResetToken"
// @Param new_password query string false "NewPassword"
// @Success 200 {object} string "Password reset successfully"
// @Failure 400 {object} string "Invalid input"
// @Failure 500 {object} string "Internal server error"
// @Security BearerAuth
// @Router /reset-password [post]
func (h *Handler) ResetPassword(c *gin.Context) {
	var resetCode ap.UsersResetPassword
	reset_token := c.Query("reset_token")
	new_password := c.Query("new_password")

	if new_password == "" || new_password == "string" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Passwrod is empty"})
	}

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	email := claims.(jwt.MapClaims)["email"].(string)

	resetCode.NewPassword = new_password
	resetCode.ResetToken = reset_token
	resetCode.Email = email

	_, err := h.User.ResetPassword(context.Background(), &resetCode)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Password reset successfully"})
}

// RefreshToken godoc
// @Summary Get token
// @Description Get the token of the authenticated user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} ap.TokenRes
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "User not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /refresh-token [get]
func (h *Handler) RefreshToken(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := claims.(jwt.MapClaims)["user_id"].(string)

	token, err := h.User.RefreshToken(context.Background(), &ap.ById{Id: id})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, token)
}
