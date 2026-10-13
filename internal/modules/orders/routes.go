package orders

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	orderRoutes := r.Group("/orders")
	{
		orderRoutes.POST("/checkout", handler.Checkout)
		orderRoutes.GET("", handler.GetOrders)
		orderRoutes.GET("/:id", handler.GetOrderByID)
	}
}
