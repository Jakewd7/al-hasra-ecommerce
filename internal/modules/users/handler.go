package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// Refresh Token godoc
// @Tags user
// @Router /users/refresh [post]
func (h *Handler) Refresh(c *gin.Context) {
	var input RefreshRequest

	cookieToken, err := c.Cookie("refresh_token")
	if err == nil {
		input.RefreshToken = cookieToken
	} else {
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "refresh_token required"})
			return
		}
	}

	response, err := h.service.RefreshToken(input.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid login session"})
		return
	}

	c.SetCookie("access_token", response.AccessToken, 900, "/", "", false, true)
	c.SetCookie("refresh_token", response.RefreshToken, 3600*24*7, "/", "", false, true)

	c.JSON(http.StatusOK, response)
}

// Register godoc
// @Summary Register new user
// @Description Register a new user account with role
// @Tags Users
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Register Request"
// @Success 201 {object} UserResponse
// @Failure 400 {object} map[string]string
// @Router /users/register [post]
func (h *Handler) Register(c *gin.Context) {
	var input RegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser, err := h.service.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

// Login godoc
// @Summary Login user
// @Description Login and generate Access Token
// @Tags Users
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login Request"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} map[string]string
// @Router /users/login [post]
func (h *Handler) Login(c *gin.Context) {
	var input LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.service.LoginUser(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("access_token", response.AccessToken, 3600*24, "/", "", false, true)
	c.SetCookie("refresh_token", response.RefreshToken, 3600*24*7, "/", "", false, true)

	c.JSON(http.StatusOK, response)
}

// Logout godoc
// @Summary Logout user
// @Description Clear authentication cookie
// @Tags Users
// @Success 200 {object} map[string]string
// @Router /users/logout [post]
func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

// GetUsers godoc
// @Summary Get all users
// @Description Get list of all users
// @Tags Users
// @Produce json
// @Success 200 {array} UserResponse
// @Router /users [get]
func (h *Handler) GetUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// UpdateUser godoc
// @Summary Update user
// @Description Update user details including Role
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body UpdateUserRequest true "Update Request"
// @Success 200 {object} UserResponse
// @Failure 400 {object} map[string]string
// @Router /users/{id} [put]
func (h *Handler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter is required"})
		return
	}

	var input UpdateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := h.service.UpdateUser(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser godoc
// @Summary Delete user
// @Description Soft delete user by ID
// @Tags Users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]string
// @Router /users/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter is required"})
		return
	}

	if err := h.service.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
