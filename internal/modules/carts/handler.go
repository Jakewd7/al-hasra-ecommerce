package carts

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

// AddToCart godoc
// @Summary Add product to cart
// @Description Add product item to cart (guest mode or authenticated)
// @Tags Carts
// @Accept json
// @Produce json
// @Param request body AddToCartRequest true "Add To Cart Request"
// @Success 201 {object} CartResponse
// @Failure 400 {object} map[string]string
// @Router /carts [post]
func (h *Handler) AddToCart(c *gin.Context) {
	var input AddToCartRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cart, err := h.service.AddToCart(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, cart)
}

// GetCarts godoc
// @Summary Get all carts
// @Description Get list of cart items
// @Tags Carts
// @Produce json
// @Success 200 {array} CartResponse
// @Router /carts [get]
func (h *Handler) GetCarts(c *gin.Context) {
	carts, err := h.service.GetAllCarts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, carts)
}

// DeleteCart godoc
// @Summary Delete cart item
// @Description Delete cart item by ID
// @Tags Carts
// @Param id path string true "Cart ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /carts/{id} [delete]
func (h *Handler) DeleteCart(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter is required"})
		return
	}

	if err := h.service.DeleteCart(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cart item deleted successfully"})
}
