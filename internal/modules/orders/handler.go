package orders

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

// Checkout godoc
// @Summary Process order checkout
// @Description Process checkout without requiring user authentication (Guest Checkout)
// @Tags Orders
// @Accept json
// @Produce json
// @Param request body CheckoutRequest true "Checkout Request"
// @Success 201 {object} OrderResponse
// @Failure 400 {object} map[string]string
// @Router /orders/checkout [post]
func (h *Handler) Checkout(c *gin.Context) {
	var input CheckoutRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.service.Checkout(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, order)
}

// GetOrders godoc
// @Summary Get all orders
// @Description Get list of all submitted orders
// @Tags Orders
// @Produce json
// @Success 200 {array} OrderResponse
// @Router /orders [get]
func (h *Handler) GetOrders(c *gin.Context) {
	orders, err := h.service.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// GetOrderByID godoc
// @Summary Get order by ID
// @Description Get order details by ID
// @Tags Orders
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} OrderResponse
// @Failure 404 {object} map[string]string
// @Router /orders/{id} [get]
func (h *Handler) GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	order, err := h.service.GetOrderByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}
